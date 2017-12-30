/*
Package gdax implements a client library for the gdax api. See https://docs.gdax.com/ for api documentation.
*/
package gdax

import (
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

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
		privateLimiter: rate.NewLimiter(rate.Every(time.Second/5), 1),
		publicLimiter:  rate.NewLimiter(rate.Every(time.Second/3), 1),
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
		privateLimiter: rate.NewLimiter(rate.Every(time.Second), 5),
		publicLimiter:  rate.NewLimiter(rate.Every(time.Second), 3),
		httpClient:     httpClient,
	}

	return &client
}
