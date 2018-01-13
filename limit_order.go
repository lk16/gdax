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
	Size      decimal.Decimal `json:"size,string"`
	Side                      `json:"side"`
	ProductId string          `json:"product_id"`
	SelfTradePrevention       `json:"stp,omitempty"`

	Price      decimal.Decimal `json:"price,string"`
	TimeInForce                `json:"time_in_force"`
	ExpireTime *Time           `json:"expire_time,string,omitempty"`
	PostOnly   bool            `json:"post_only"`

	ID            uuid.UUID       `json:"id"`
	Status        OrderStatus     `json:"status"`
	Settled       bool            `json:"settled"`
	CreatedAt     Time            `json:"created_at,string"`
	FillFees      decimal.Decimal `json:"fill_fees,string"`
	FilledSize    decimal.Decimal `json:"filled_size,string"`
	ExecutedValue decimal.Decimal `json:"executed_value,string"`
}

func (c *Client) CreateLimitOrder(ctx context.Context, newOrder *LimitOrderRequest) (LimitOrderResponse, error) {
	newOrder.Type = "limit"
	var response LimitOrderResponse
	url := fmt.Sprintf("/orders")
	_, err := c.request(ctx, true, "POST", url, newOrder, &response)
	return response, err
}
