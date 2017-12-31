package gdax

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

type Currency struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	MinSize decimal.Decimal `json:"min_size,string"`
}

func (c *Client) GetCurrencies(ctx context.Context) ([]Currency, error) {
	var currencies []Currency

	url := fmt.Sprintf("/currencies")
	_, err := c.request(ctx, false, "GET", url, nil, &currencies)
	return currencies, err
}
