package gdax

import (
	"context"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CoinbaseAccountType string

const (
	CoinbaseAccountTypeWallet CoinbaseAccountType = "wallet"
	CoinbaseAccountTypeFiat CoinbaseAccountType = "fiat"
)

type CoinbaseAccount struct {
	ID       uuid.UUID           `json:"id,string"`
	Name     string              `json:"name,string"`
	Balance  decimal.Decimal     `json:"balance,string"`
	Currency string              `json:"currency, string"`
	Type     CoinbaseAccountType `json:"type,string"`
	Primary  bool                `json:"primary,bool"`
	Active   bool                `json:"active,bool"`

	// skipping the account information
}

func (c *Client) GetCoinbaseAccounts(ctx context.Context) ([]CoinbaseAccount, error) {
	var accounts []CoinbaseAccount
	_, err := c.request(ctx, true, "GET", "/coinbase-accounts", nil, &accounts)
	return accounts, err
}
