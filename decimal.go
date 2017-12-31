package gdax

import "github.com/shopspring/decimal"

func requireDecimalFromString(s string) decimal.Decimal {
	dec, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return dec
}
