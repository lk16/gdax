package gdax

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

type Fill struct {
	TradeId   int             `json:"trade_id,int"`
	ProductId string          `json:"product_id"`
	Price     decimal.Decimal `json:"price,string"`
	Size      decimal.Decimal `json:"size,string"`
	FillId    string          `json:"order_id"`
	CreatedAt Time            `json:"created_at,string"`
	Fee       decimal.Decimal `json:"fee,string"`
	Settled   bool            `json:"settled"`
	Side      string          `json:"side"`
	Liquidity string          `json:"liquidity"`
}

type ListFillsParams struct {
	OrderId    string
	ProductId  string
	Pagination PaginationParams
}

func (c *Client) ListFills(ctx context.Context, p ...ListFillsParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
		if p[0].OrderId != "" {
			paginationParams.AddExtraParam("order_id", p[0].OrderId)
		}
		if p[0].ProductId != "" {
			paginationParams.AddExtraParam("product_id", p[0].ProductId)
		}
	}

	return NewCursor(ctx, c, true, "GET", fmt.Sprintf("/fills"),
		&paginationParams)
}
