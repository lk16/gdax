package gdax

import (
	"context"
	"fmt"
	//"github.com/google/uuid"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderType string

const (
	Limit  OrderType = "limit"
	Market OrderType = "market"
	Stop   OrderType = "stop"
)

type Side string

const (
	Buy  Side = "buy"
	Sell Side = "sell"
)

type SelfTradePrevention string

const (
	DecrementAndCancel SelfTradePrevention = "dc"
	CancelOldest       SelfTradePrevention = "co"
	CancelNewest       SelfTradePrevention = "cn"
	CancelBoth         SelfTradePrevention = "cb"
)

type TimeInForce string

const (
	GoodTillCanceled  TimeInForce = "GTC"
	GoodTillTime      TimeInForce = "GTT"
	ImmediateOrCancel TimeInForce = "IOC"
	FillOrKill        TimeInForce = "FOK"
)

type CancelAfter string

const (
	CancelAfterMin  CancelAfter = "min"
	CancelAfterHour CancelAfter = "hour"
	CancelAfterDay  CancelAfter = "day"
)

type OrderStatus string

const (
	OrderStatusPending OrderStatus = "pending"
	OrderStatusOpen    OrderStatus = "open"
	OrderStatusActive  OrderStatus = "active"
	OrderStatusDone    OrderStatus = "done"
)

type OrderRequest struct {
	Type                OrderType       `json:"type"`
	Size                decimal.Decimal `json:"size,string,omitempty"`
	Side                `json:"side"`
	ProductId           string `json:"product_id"`
	ClientOID           string `json:"client_oid,omitempty"`
	SelfTradePrevention `json:"stp,omitempty"`
	// Limit OrderRequest
	Price       decimal.Decimal `json:"price,string,omitempty"`
	TimeInForce `json:"time_in_force,omitempty"`
	CancelAfter CancelAfter `json:"cancel_after,omitempty"`
	PostOnly    bool        `json:"post_only,omitempty"`
	// Market OrderRequest
	Funds string `json:"funds,omitempty"`
}

type OrderResponse struct {
	Type                OrderType       `json:"type"`
	Size                decimal.Decimal `json:"size,string,omitempty"`
	Side                `json:"side"`
	ProductId           string `json:"product_id"`
	ClientOID           string `json:"client_oid,omitempty"`
	SelfTradePrevention `json:"stp,omitempty"`
	// Limit OrderRequest
	Price       decimal.Decimal `json:"price,string,omitempty"`
	TimeInForce `json:"time_in_force,omitempty"`
	CancelAfter CancelAfter `json:"cancel_after,omitempty"`
	PostOnly    bool        `json:"post_only,omitempty"`
	// Market OrderRequest
	Funds decimal.Decimal `json:"funds,string,omitempty"`
	// Response Fields
	ID             uuid.UUID       `json:"id,omitempty"`
	Status         OrderStatus     `json:"status,omitempty"`
	Settled        bool            `json:"settled,omitempty"`
	DoneReason     string          `json:"done_reason,omitempty"`
	CreatedAt      Time            `json:"created_at,string,omitempty"`
	SpecifiedFunds decimal.Decimal `json:"specified_funds,omitempty"`
	FillFees       decimal.Decimal `json:"fill_fees,string,omitempty"`
	FilledSize     decimal.Decimal `json:"filled_size,string,omitempty"`
	ExecutedValue  decimal.Decimal `json:"executed_value,string,omitempty"`
}

type CancelAllOrdersParams struct {
	ProductId string
}

type ListOrdersParams struct {
	Status     string
	ProductId  string
	Pagination PaginationParams
}

func (c *Client) CreateOrder(ctx context.Context, newOrder *OrderRequest) (OrderResponse, error) {
	var savedOrder OrderResponse

	if len(newOrder.Type) == 0 {
		newOrder.Type = "limit"
	}

	url := fmt.Sprintf("/orders")
	_, err := c.request(ctx, true, "POST", url, newOrder, &savedOrder)
	return savedOrder, err
}

func (c *Client) CancelOrder(ctx context.Context, id uuid.UUID) error {
	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.request(ctx, true, "DELETE", url, nil, nil)
	return err
}

func (c *Client) CancelAllOrders(ctx context.Context, p ...CancelAllOrdersParams) ([]string, error) {
	var orderIDs []string
	url := "/orders"

	if len(p) > 0 && p[0].ProductId != "" {
		url = fmt.Sprintf("%s?product_id=", p[0].ProductId)
	}

	_, err := c.request(ctx, true, "DELETE", url, nil, &orderIDs)
	return orderIDs, err
}

func (c *Client) GetOrder(ctx context.Context, id uuid.UUID) (OrderResponse, error) {
	var savedOrder OrderResponse

	url := fmt.Sprintf("/orders/%s", id)
	_, err := c.request(ctx, true, "GET", url, nil, &savedOrder)
	return savedOrder, err
}

func (c *Client) ListOrders(ctx context.Context, p ...ListOrdersParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
		if p[0].Status != "" {
			paginationParams.AddExtraParam("status", p[0].Status)
		}
		if p[0].ProductId != "" {
			paginationParams.AddExtraParam("product_id", p[0].ProductId)
		}
	}

	return NewCursor(ctx,
		c,
		true,
		"GET",
		fmt.Sprintf("/orders"),
		&paginationParams,
	)
}
