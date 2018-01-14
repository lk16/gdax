package gdax

import (
	"github.com/shopspring/decimal"
	"fmt"
	"context"
	"github.com/google/uuid"
	"errors"
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
	SelfTradePrevention       `json:"stp,omitempty"`

	Size     *decimal.Decimal `json:"size,string,omitempty"`
	Funds    *decimal.Decimal `json:"funds,string,omitempty"`
	PostOnly bool             `json:"post_only"`

	ID            uuid.UUID       `json:"id"`
	Status        OrderStatus     `json:"status,omitempty"`
	Settled       bool            `json:"settled"`
	CreatedAt     Time            `json:"created_at,string,omitempty"`
	FillFees      decimal.Decimal `json:"fill_fees,string,omitempty"`
	FilledSize    decimal.Decimal `json:"filled_size,string,omitempty"`
	ExecutedValue decimal.Decimal `json:"executed_value,string,omitempty"`
}

func (c *Client) CreateStopOrder(ctx context.Context, newOrder *StopOrderRequest) (StopOrderResponse, error) {
	var response StopOrderResponse
	newOrder.Type = "stop"

	// To prevent confusion.
	if newOrder.Side == Sell && newOrder.Funds != nil && newOrder.Funds.GreaterThan(decimal.Zero) {
		return response, errors.New("sell-stop orders (stop-loss) should specify Size and not Funds")
	} else if newOrder.Side == Buy && newOrder.Size != nil && newOrder.Size.GreaterThan(decimal.Zero) {
		return response, errors.New("buy-stop orders should specify Funds and not Size")
	}

	url := fmt.Sprintf("/orders")
	_, err := c.request(ctx, true, "POST", url, newOrder, &response)
	return response, err
}
