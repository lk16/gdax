Go GDAX [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/randyp/gdax) [![Build Status](https://travis-ci.org/randyp/gdax.svg?branch=master)](https://travis-ci.org/randyp/gdax) [![Go Report Card](https://goreportcard.com/badge/github.com/randyp/gdax)](https://goreportcard.com/report/github.com/randyp/gdax)
========

## Summary

Go client library for [GDAX](https://www.gdax.com) restful endpoints and websocket feed. A heavily modified fork [github.com/preichenberger/go-gdax](//github.com/preichenberger/go-gdax) so that we could:
* use shopspring decimals for currenty amounts, since floats will occasionally be too precise
* use uuid.UUID for all possible fields
* add missing fields and optional query parameters
* support rate limiting
* support public-endpoint-only clients
* adhere to go coding hygenic standards

## Installation

```sh
go get github.com/randy/go-gdax
```

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
