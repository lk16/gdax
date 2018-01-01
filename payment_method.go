package gdax

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"context"
)

type PaymentMethodAmount struct {
	Amount   decimal.Decimal `json:"amount,string"`
	Currency string          `json:"currency"`
}

type PaymentMethodLimit struct {
	PeriodInDays int                 `json:"period_in_days"`
	Total        PaymentMethodAmount `json:"total"`
	Remaining    PaymentMethodAmount `json:"remaining"`
}

type PaymentMethod struct {
	ID            uuid.UUID                       `json:"id,string"`
	Type          string                          `json:"type"`
	Name          string                          `json:"name"`
	Currency      string                          `json:"currency"`
	PrimaryBuy    bool                            `json:"primary_buy"`
	PrimarySell   bool                            `json:"primary_sell"`
	AllowBuy      bool                            `json:"allow_buy"`
	AllowSell     bool                            `json:"allow_sell"`
	AllowDeposit  bool                            `json:"allow_deposit"`
	AllowWithdraw bool                            `json:"allow_withdraw"`
	Limits        map[string][]PaymentMethodLimit `json:"limits"`
}

func (c *Client) GetPaymentMethods(ctx context.Context) ([]PaymentMethod, error) {
	var paymentMethods []PaymentMethod
	_, err := c.request(ctx, true, "GET", "/payment-methods", nil, &paymentMethods)
	return paymentMethods, err
}