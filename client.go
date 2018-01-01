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

type Client struct {
	BaseURL        string
	secret         string
	key            string
	passphrase     string
	hasCredentials bool
	privateLimiter *rate.Limiter
	publicLimiter  *rate.Limiter
	httpClient     *http.Client
}

// NewPublicClient creates a client for the public restful endpoint only.
// Calling any methods that use the private restful endpoint will return errors.
// Requests to the public endpoint are rate limited to 3 requests per second.
//
// Example:
// import (
//   "context"
//   "github.com/randyp/gdax"
//   "net/http"
//  )
//
// client := NewPublicClient(&http.Client{
//   Timeout: 15 * time.Second,
// })
// currencies := client.GetCurrencies(context.Background())
// // do something with the currencies
func NewPublicClient(httpClient *http.Client) *Client {
	client := Client{
		BaseURL:        "https://api.gdax.com",
		secret:         "",
		key:            "",
		passphrase:     "",
		hasCredentials: false,
		privateLimiter: rate.NewLimiter(rate.Every(200*time.Millisecond), 1),
		publicLimiter:  rate.NewLimiter(rate.Every(333*time.Millisecond), 1),
		httpClient:     httpClient,
	}

	return &client
}

// NewClient creates a client for the public and private restful endpoints.
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
//   os.Getenv("COINBASE_SECRET"),
//   os.Getenv("COINBASE_KEY"),
//   os.Getenv("COINBASE_PASSPHRASE"),
// )
// accounts := client.GetAccounts(context.Background())
// // do something with the accounts
func NewClient(httpClient *http.Client, secret, key, passphrase string) *Client {
	client := Client{
		BaseURL:        "https://api.gdax.com",
		secret:         secret,
		key:            key,
		passphrase:     passphrase,
		hasCredentials: true,
		privateLimiter: rate.NewLimiter(rate.Every(200*time.Millisecond), 1),
		publicLimiter:  rate.NewLimiter(rate.Every(333*time.Millisecond), 1),
		httpClient:     httpClient,
	}

	return &client
}

func (c *Client) request(ctx context.Context, private bool, method string, url string, params, result interface{}) (res *http.Response, err error) {
	if private && !c.hasCredentials {
		return res, errors.New("cannot use a public client to make requests to the private api")
	}

	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
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

	if c.hasCredentials {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		h, err := c.Headers(method, url, timestamp, string(data))
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

	res, err = c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		defer res.Body.Close()
		coinbaseError := Error{}
		decoder := json.NewDecoder(res.Body)
		if err := decoder.Decode(&coinbaseError); err != nil {
			return res, err
		}

		return res, error(coinbaseError)
	}

	if result != nil {
		//jsonBody, err := ioutil.ReadAll(res.Body)
		//decoder := json.NewDecoder(bytes.NewReader(jsonBody))
		//fmt.Printf("%s\n", string(jsonBody))

		decoder := json.NewDecoder(res.Body)
		if err = decoder.Decode(result); err != nil {
			return res, err
		}
	}

	return res, nil
}

// Headers generates a map that can be used as headers to authenticate a request
func (c *Client) Headers(method, url, timestamp, data string) (map[string]string, error) {
	h := make(map[string]string)
	h["CB-ACCESS-KEY"] = c.key
	h["CB-ACCESS-PASSPHRASE"] = c.passphrase
	h["CB-ACCESS-TIMESTAMP"] = timestamp

	message := fmt.Sprintf(
		"%s%s%s%s",
		timestamp,
		method,
		url,
		data,
	)

	sig, err := c.generateSig(message, c.secret)
	if err != nil {
		return nil, err
	}
	h["CB-ACCESS-SIGN"] = sig
	return h, nil
}

func (c *Client) generateSig(message, secret string) (string, error) {
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
