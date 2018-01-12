package gdax

import (
	"github.com/shopspring/decimal"
	"fmt"
	"context"
)

type StopOrderRequest struct {
	// Ignore this
	Type OrderType `json:"type"`

	// Required
	Side      Side            `json:"side"`
	ProductId string          `json:"product_id"`
	Price     decimal.Decimal `json:"price,string"`
	Size      decimal.Decimal `json:"size,string"`
	Funds      decimal.Decimal `json:"funds,string"`

	// Optional
	ClientOID string    `json:"client_oid,omitempty"`
	SelfTradePrevention `json:"stp,omitempty"`
}


type StopOrderResponse struct {
	Type      OrderType       `json:"type"`
	Size      decimal.Decimal `json:"size,string,omitempty"`
	Side                      `json:"side"`
	ProductId string          `json:"product_id"`
	ClientOID string          `json:"client_oid,omitempty"`
	SelfTradePrevention       `json:"stp,omitempty"`

	Price       decimal.Decimal `json:"price,string,omitempty"`
	TimeInForce                 `json:"time_in_force,omitempty"`
	CancelAfter CancelAfter     `json:"cancel_after,omitempty"`
	PostOnly    bool            `json:"post_only"`

	ID            string          `json:"id,omitempty"`
	Status        string          `json:"status,omitempty"`
	Settled       bool            `json:"settled"`
	DoneReason    string          `json:"done_reason,omitempty"`
	CreatedAt     Time            `json:"created_at,string,omitempty"`
	FillFees      decimal.Decimal `json:"fill_fees,string,omitempty"`
	FilledSize    decimal.Decimal `json:"filled_size,string,omitempty"`
	ExecutedValue decimal.Decimal `json:"executed_value,string,omitempty"`
}

func (c *Client) CreateStopOrder(ctx context.Context, newOrder *StopOrderRequest) (StopOrderResponse, error) {
	newOrder.Type = "stop"
	var response StopOrderResponse
	url := fmt.Sprintf("/orders")
	_, err := c.request(ctx, true, "POST", url, newOrder, &response)
	return response, err
}
