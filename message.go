package gdax

import (
	"github.com/shopspring/decimal"
)

type Message struct {
	Type          string              `json:"type"`
	ProductId     string              `json:"product_id,omitempty"`
	ProductIds    []string            `json:"product_ids,omitempty"`
	TradeId       int                 `json:"trade_id,number,omitempty"`
	OrderId       string              `json:"order_id,omitempty"`
	Sequence      int64               `json:"sequence,numbe,omitemptyr"`
	MakerOrderId  string              `json:"maker_order_id,omitempty"`
	TakerOrderId  string              `json:"taker_order_id,omitempty"`
	Time          Time                `json:"time,string"`
	RemainingSize decimal.Decimal     `json:"remaining_size,string"`
	NewSize       decimal.Decimal     `json:"new_size,string"`
	OldSize       decimal.Decimal     `json:"old_size,string"`
	Size          decimal.Decimal     `json:"size,string"`
	Price         decimal.Decimal     `json:"price,string"`
	Side          string              `json:"side"`
	Reason        string              `json:"reason"`
	OrderType     string              `json:"order_type"`
	Funds         decimal.Decimal     `json:"funds,string"`
	NewFunds      decimal.Decimal     `json:"new_funds,string"`
	OldFunds      decimal.Decimal     `json:"old_funds,string"`
	Message       string              `json:"message"`
	Bids          [][]decimal.Decimal `json:"bids,omitempty"`
	Asks          [][]decimal.Decimal `json:"asks,omitempty"`
	Changes       [][]string          `json:"changes,omitempty"`
	LastSize      decimal.Decimal     `json:"last_size,string"`
	BestBid       decimal.Decimal     `json:"best_bid,string"`
	BestAsk       decimal.Decimal     `json:"best_ask,string"`
	Channels      []MessageChannel    `json:"channels"`

	// For User channel
	UserId    string `json:"user_id,omitempty"`
	ProfileId string `json:"profile_id,omitempty"`

	// Authentication
	Signature  string `json:"signature,omitempty"`
	Key        string `json:"key,omitempty"`
	Passphrase string `json:"passphrase,omitempty"`
	Timestamp  string `json:"timestamp,omitempty"`
}

type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}
