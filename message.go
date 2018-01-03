package gdax

import "github.com/shopspring/decimal"

type Message struct {
	Type          string           `json:"type"`
	ProductId     string           `json:"product_id"`
	ProductIds    []string         `json:"product_ids"`
	TradeId       int              `json:"trade_id,number"`
	OrderId       string           `json:"order_id"`
	Sequence      int64            `json:"sequence,number"`
	MakerOrderId  string           `json:"maker_order_id"`
	TakerOrderId  string           `json:"taker_order_id"`
	Time          Time             `json:"time,string"`
	RemainingSize decimal.Decimal  `json:"remaining_size,string"`
	NewSize       decimal.Decimal  `json:"new_size,string"`
	OldSize       decimal.Decimal  `json:"old_size,string"`
	Size          decimal.Decimal  `json:"size,string"`
	Price         decimal.Decimal  `json:"price,string"`
	Side          string           `json:"side"`
	Reason        string           `json:"reason"`
	OrderType     string           `json:"order_type"`
	Funds         decimal.Decimal  `json:"funds,string"`
	NewFunds      decimal.Decimal  `json:"new_funds,string"`
	OldFunds      decimal.Decimal  `json:"old_funds,string"`
	Message       string           `json:"message"`
	Bids          [][]string       `json:"bids,omitempty"`
	Asks          [][]string       `json:"asks,omitempty"`
	Changes       [][]string       `json:"changes,omitempty"`
	LastSize      decimal.Decimal  `json:"last_size,string"`
	BestBid       decimal.Decimal  `json:"best_bid,string"`
	BestAsk       decimal.Decimal  `json:"best_ask,string"`
	Channels      []MessageChannel `json:"channels"`
}

type MessageChannel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}
