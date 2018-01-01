package gdax

import (
	"github.com/shopspring/decimal"
	"context"
)

type TrailingVolume struct {
	ProductID      string          `json:"product_id"`
	ExchangeVolume decimal.Decimal `json:"exchange_volume,string"`
	Volume         decimal.Decimal `json:"volume,string"`
	RecordedAt     Time            `json:"recorded_at,string"`
}

func (c *Client) GetTrailingVolume(ctx context.Context) ([]TrailingVolume, error) {
	var volume []TrailingVolume
	_, err := c.request(ctx, true, "GET", "/users/self/trailing-volume", nil, &volume)
	return volume, err
}
