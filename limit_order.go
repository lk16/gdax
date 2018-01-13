package gdax

import (
	"github.com/shopspring/decimal"
	"fmt"
	"context"
	"github.com/google/uuid"
)

type LimitOrderRequest struct {
	// Ignore this
	Type OrderType `json:"type"`

	// Required
	Side      Side            `json:"side"`
	ProductId string          `json:"product_id"`
	Price     decimal.Decimal `json:"price,string"`
	Size      decimal.Decimal `json:"size,string"`

	// Optional
	ClientOID *uuid.UUID `json:"client_oid,omitempty"`
	SelfTradePrevention  `json:"stp,omitempty"`
	TimeInForce          `json:"time_in_force,omitempty"`
	CancelAfter          `json:"cancel_after,omitempty"`
	PostOnly  *bool      `json:"post_only,omitempty"`
}

type LimitOrderResponse struct {
	Type      OrderType       `json:"type"`
	Size      decimal.Decimal `json:"size,string,omitempty"`
	Side                      `json:"side"`
	ProductId string          `json:"product_id"`
	SelfTradePrevention       `json:"stp,omitempty"`

	Price      decimal.Decimal `json:"price,string,omitempty"`
	TimeInForce                `json:"time_in_force,omitempty"`
	ExpireTime *Time           `json:"expire_time,string,omitempty"`
	PostOnly   bool            `json:"post_only"`

	ID            uuid.UUID       `json:"id,omitempty"`
	Status        OrderStatus     `json:"status,omitempty"`
	Settled       bool            `json:"settled"`
	DoneReason    string          `json:"done_reason,omitempty"`
	CreatedAt     Time            `json:"created_at,string,omitempty"`
	FillFees      decimal.Decimal `json:"fill_fees,string,omitempty"`
	FilledSize    decimal.Decimal `json:"filled_size,string,omitempty"`
	ExecutedValue decimal.Decimal `json:"executed_value,string,omitempty"`
}

func (c *Client) CreateLimitOrder(ctx context.Context, newOrder *LimitOrderRequest) (LimitOrderResponse, error) {
	newOrder.Type = "limit"
	var response LimitOrderResponse
	url := fmt.Sprintf("/orders")
	_, err := c.request(ctx, true, "POST", url, newOrder, &response)
	return response, err
}
