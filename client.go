/*
Package gdax implements a client library for the gdax api. See https://docs.gdax.com/ for api documentation.
*/
package gdax

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/time/rate"
	"net/http"
	"strconv"
	"time"
)

type Authentication struct {
	Secret     string
	Key        string
	Passphrase string
}

type Client struct {
	HttpClient     *http.Client
	BaseURL        string
	Authentication *Authentication
	privateLimiter *rate.Limiter
	publicLimiter  *rate.Limiter
}

type FeedAuthHeaders struct {
	Signature  string
	Key        string
	Passphrase string
	Timestamp  string
}

// NewClient creates a client for the public endpoint and optionally the private endpoints if a non-nil valid Authentication is passed.
// Requests to the public endpoint are rate limited to 3 requests per second.
// Requests to the private endpoint are rate limited to 5 requests per second.
//
// Example:
// import (
//   "context"
//   "github.com/randyp/gdax"
//   "net/http"
//   "os"
//  )
//
// client := NewClient(
// 	 &http.Client{
//     Timeout: 15 * time.Second,
//   },
//   &gdax.Authentication{
//     Secret: os.Getenv("COINBASE_SECRET"),
//     Key: os.Getenv("COINBASE_KEY"),
//     Passphrase: os.Getenv("COINBASE_PASSPHRASE"),
//	 }
// )
// accounts := client.GetAccounts(context.Background())
// // do something with the accounts
func NewClient(httpClient *http.Client, authentication *Authentication) *Client {
	client := Client{
		BaseURL:        "https://api.gdax.com",
		Authentication: authentication,
		privateLimiter: rate.NewLimiter(rate.Every(200*time.Millisecond), 1),
		publicLimiter:  rate.NewLimiter(rate.Every(333*time.Millisecond), 1),
		HttpClient:     httpClient,
	}

	return &client
}

func (c *Client) request(ctx context.Context, private bool, method string, url string, params, result interface{}) (res *http.Response, err error) {
	if private && c.Authentication == nil {
		return res, errors.New("cannot use a public client to make requests to the private api")
	}

	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}

		if _, ok := params.(*OrderRequest); ok {
			var dataObj map[string]interface{}

			json.Unmarshal(data, &dataObj)

			f, err := strconv.ParseFloat(dataObj["size"].(string), 64)
			if err != nil {
				return res, nil
			}
			if f == 0.0 {
				delete(dataObj, "size")
			}
			data, err = json.Marshal(dataObj)
			if err != nil {
				return res, nil
			}
		}

		body = bytes.NewReader(data)
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, url)
	req, err := http.NewRequest(method, fullURL, body)
	if err != nil {
		return res, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "github.com/randy/gdax 0.0.1")

	if c.Authentication != nil {
		h, err := c.Authentication.RequestHeaders(method, url, string(data))
		if err != nil {
			return res, err
		}
		for k, v := range h {
			req.Header.Add(k, v)
		}
	}

	if private {
		c.privateLimiter.Wait(ctx)
	} else {
		c.publicLimiter.Wait(ctx)
	}

	res, err = c.HttpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		defer res.Body.Close()
		coinbaseError := Error{}

		//body, _ := ioutil.ReadAll(res.Body)
		//log.Println(string(body))
		//decoder := json.NewDecoder(bytes.NewReader(body))

		decoder := json.NewDecoder(res.Body)

		if err := decoder.Decode(&coinbaseError); err != nil {
			return res, err
		}

		return res, error(coinbaseError)
	}

	if result != nil {
		//body, _ := ioutil.ReadAll(res.Body)
		//log.Println(string(body))
		//decoder := json.NewDecoder(bytes.NewReader(body))

		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}

// Headers generates a map that can be used as headers to authenticate a request
func (a *Authentication) RequestHeaders(method, url, data string) (map[string]string, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	h := make(map[string]string)
	h["CB-ACCESS-KEY"] = a.Key
	h["CB-ACCESS-PASSPHRASE"] = a.Passphrase
	h["CB-ACCESS-TIMESTAMP"] = timestamp

	message := fmt.Sprintf(
		"%s%s%s%s",
		timestamp,
		method,
		url,
		data,
	)

	sig, err := a.generateSig(message, a.Secret)
	if err != nil {
		return nil, err
	}
	h["CB-ACCESS-SIGN"] = sig
	return h, nil
}

func (a *Authentication) NewFeedAuthHeaders() (h FeedAuthHeaders, err error) {
	authHeaders, err := a.RequestHeaders("GET", "/users/self/verify", "")

	h.Signature = authHeaders["CB-ACCESS-SIGN"]
	h.Key = authHeaders["CB-ACCESS-KEY"]
	h.Passphrase = authHeaders["CB-ACCESS-PASSPHRASE"]
	h.Timestamp = authHeaders["CB-ACCESS-TIMESTAMP"]
	return h, err
}

func (a *Authentication) generateSig(message, secret string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	signature := hmac.New(sha256.New, key)
	_, err = signature.Write([]byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature.Sum(nil)), nil
}
