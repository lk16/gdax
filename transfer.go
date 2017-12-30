package gdax

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

type Transfer struct {
	Type              string          `json:"type"`
	Amount            decimal.Decimal `json:"amount,string"`
	CoinbaseAccountId string          `json:"coinbase_account_id,string"`
}

func (c *Client) CreateTransfer(ctx context.Context, newTransfer *Transfer) (Transfer, error) {
	var savedTransfer Transfer

	url := fmt.Sprintf("/transfers")
	_, err := c.request(ctx, true, "POST", url, newTransfer, &savedTransfer)
	return savedTransfer, err
}
