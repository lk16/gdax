package gdax

import (
	"github.com/shopspring/decimal"
	"fmt"
	"context"
	"github.com/google/uuid"
)

type StopOrderRequest struct {
	// Ignore this
	Type OrderType `json:"type"`

	// Required
	Side      Side            `json:"side"`
	ProductId string          `json:"product_id"`
	Price     decimal.Decimal `json:"price,string"`

	// Pick Size Or Funds
	Size  *decimal.Decimal `json:"size,string,omitempty"`
	Funds *decimal.Decimal `json:"funds,string,omitempty"`

	// Optional
	ClientOID *uuid.UUID `json:"client_oid,omitempty"`
	SelfTradePrevention  `json:"stp,omitempty"`
}

type StopOrderResponse struct {
	Type      OrderType       `json:"type"`
	Side                      `json:"side"`
	ProductId string          `json:"product_id"`
	Price     decimal.Decimal `json:"price,string"`
	Size      *decimal.Decimal `json:"size,string,omitempty"`
	Funds     *decimal.Decimal `json:"funds,string,omitempty"`
	SelfTradePrevention       `json:"stp,omitempty"`

	PostOnly    bool        `json:"post_only"`

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
