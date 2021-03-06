Go GDAX [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/randyp/gdax) [![Build Status](https://travis-ci.org/randyp/gdax.svg?branch=master)](https://travis-ci.org/randyp/gdax) [![Go Report Card](https://goreportcard.com/badge/github.com/randyp/gdax)](https://goreportcard.com/report/github.com/randyp/gdax)
========

## Installation
```sh
go get -u github.com/randyp/gdax
```

### Summary
Go client library for [GDAX](https://www.gdax.com) restful endpoints and websocket feed.
* shopspring decimals for currency
* use uuid.UUID where possible
* support public-endpoint-only clients
* rate limiting options


## Documentation
For full details on functionality, see [GoDoc](http://godoc.org/github.com/randy/gdax) documentation.


### Example
How to create a client and make a request:

```go
import (
  "context.Context"
  "github.com/randy/gdax"
  "os"
)

secret := os.Getenv("COINBASE_SECRET")
key := os.Getenv("COINBASE_KEY")
passphrase := os.Getenv("COINBASE_PASSPHRASE")

client := gdax.NewClient(secret, key, passphrase)
accounts := gdax.GetAccounts(context.Background())
```


### Testing
Coinbase's sandbox is down. I do not recommend trying to test against their production api. 


### Attribution
This project started as heavily modified fork [github.com/preichenberger/go-gdax](//github.com/preichenberger/go-gdax) to fix broken requests, add missing fields, use shopspring decimals, etc.