Go GDAX [![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/randyp/gdax) [![Build Status](https://travis-ci.org/randyp/gdax.svg?branch=master)](https://travis-ci.org/randyp/gdax) [![Go Report Card](https://goreportcard.com/badge/github.com/randyp/gdax)](https://goreportcard.com/report/github.com/randyp/gdax)
========

## Summary

Go client library for [GDAX](https://www.gdax.com) restful endpoints and websocket feed. A heavily modified fork github.com/preichenberger/go-gdax so that we could:
* support rate limiting
* support public-endpoint-only clients
* use shopspring decimals, since floats will occasionally not work
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

### Cursor
This library uses a cursor pattern so you don't have to keep track of pagination.

```go
cursor := client.ListOrders(context.Background())

var orders []gdax.Order
for cursor.HasMore {
  if err := cursor.NextPage(&orders); err != nil {
    println(err.Error())
    return
  }

  for _, o := range orders {
    println(o.Id)
  }
}

```

### Websockets
Listen for websocket messages

```go
  import(
    ws "github.com/gorilla/websocket"
  )

  var wsDialer ws.Dialer
  wsConn, _, err := wsDialer.Dial("wss://ws-feed.gdax.com", nil)
  if err != nil {
    println(err.Error())
  }

  subscribe := gdax.Message{
    Type:      "subscribe",
    Channels: []gdax.MessageChannel{
      gdax.MessageChannel{
        Name: "level2",
        ProductIds: []string{
          "BTC-USD",
        },
      },
    },
  }
  if err := wsConn.WriteJSON(subscribe); err != nil {
    println(err.Error())
  }

  message:= gdax.Message{}
  for true {
    if err := wsConn.ReadJSON(&message); err != nil {
      println(err.Error())
      break
    }

    if message.Type == "match" {
      println("Got a match")
    }
  }

```

### Time
Results return coinbase time type which handles different types of time parsing that GDAX returns. This wraps the native go time type

```go
  import(
    "time"
  )

  coinbaseTime := gdax.Time{}
  println(time.Time(coinbaseTime).Day())
```

### Examples
This library supports all public and private endpoints

Get Accounts:
```go
  accounts, err := client.GetAccounts()
  if err != nil {
    println(err.Error())
  }

  for _, a := range accounts {
    println(a.Balance)
  }
```

List Account Ledger:
```go
  var ledger []gdax.LedgerEntry

  accounts, err := client.GetAccounts()
  if err != nil {
    println(err.Error())
  }

  for _, a := range accounts {
    cursor := client.ListAccountLedger(a.Id)
    for cursor.HasMore {
      if err := cursor.NextPage(&ledger); err != nil {
        println(err.Error())
      }

      for _, e := range ledger {
        println(e.Amount)
      }
  }
```

Create an Order:
```go
  order := gdax.Order{
    Price: 1.00,
    Size: 1.00,
    Side: "buy",
    ProductId: "BTC-USD",
  }

  savedOrder, err := client.CreateOrder(&order)
  if err != nil {
    println(err.Error())
  }

  println(savedOrder.Id)
```

Transfer funds:
```go
  transfer := gdax.Transfer {
    Type: "deposit",
    Amount: 1.00,
  }

  savedTransfer, err := client.CreateTransfer(&transfer)
  if err != nil {
    println(err.Error())
  }
```

Get Trade history:
```go
  var trades []gdax.Trade
  cursor := client.ListTrades("BTC-USD")

  for cursor.HasMore {
    if err := cursor.NextPage(&trades); err != nil {
      for _, t := range trades {
        println(trade.CoinbaseId)
      }
    }
  }
```

### Testing
Coinbase's sandbox is down. I do not recommend trying to test against their production api. 
