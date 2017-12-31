package gdax

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID       `json:"id,string"`
	Balance   decimal.Decimal `json:"balance,string"`
	Hold      decimal.Decimal `json:"hold,string"`
	Available decimal.Decimal `json:"available,string"`
	Currency  string          `json:"currency"`
	ProfileId uuid.UUID       `json:"profile_id,string"`
}

type LedgerEntry struct {
	Id        int             `json:"id,number"`
	CreatedAt Time            `json:"created_at,string"`
	Amount    decimal.Decimal `json:"amount,string"`
	Balance   decimal.Decimal `json:"balance,string"`
	Type      string          `json:"type"`
	Details   LedgerDetails   `json:"details"`
}

type LedgerDetails struct {
	OrderId   uuid.UUID `json:"order_id"`
	TradeId   string    `json:"trade_id"`
	ProductId string    `json:"product_id"`
}

type GetAccountLedgerParams struct {
	Pagination PaginationParams
}

type Hold struct {
	Id        uuid.UUID       `json:"id,string"`
	CreatedAt Time            `json:"created_at,string"`
	UpdatedAt *Time           `json:"updated_at,string,omitempty"`
	Amount    decimal.Decimal `json:"amount,string"`
	Type      string          `json:"type"`
	Ref       uuid.UUID       `json:"ref,string"`
}

type ListHoldsParams struct {
	Pagination PaginationParams
}

func (c *Client) GetAccounts(ctx context.Context) ([]Account, error) {
	var accounts []Account
	_, err := c.request(ctx, true, "GET", "/accounts", nil, &accounts)
	return accounts, err
}

func (c *Client) GetAccount(ctx context.Context, id uuid.UUID) (Account, error) {
	account := Account{}
	url := fmt.Sprintf("/accounts/%s", id)
	_, err := c.request(ctx, true, "GET", url, nil, &account)
	return account, err
}

func (c *Client) ListAccountLedger(ctx context.Context, id uuid.UUID, p ...GetAccountLedgerParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
	}

	return NewCursor(ctx,
		c,
		true,
		"GET",
		fmt.Sprintf("/accounts/%s/ledger", id),
		&paginationParams,
	)
}

func (c *Client) ListHolds(ctx context.Context, id uuid.UUID, p ...ListHoldsParams) *Cursor {
	paginationParams := PaginationParams{}
	if len(p) > 0 {
		paginationParams = p[0].Pagination
	}

	return NewCursor(ctx,
		c,
		true,
		"GET",
		fmt.Sprintf("/accounts/%s/holds", id),
		&paginationParams,
	)
}
