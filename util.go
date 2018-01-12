package gdax

import "github.com/shopspring/decimal"

func RequireDecimalFromString(s string) decimal.Decimal {
	dec, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	return dec
}

func True() *bool {
	b := true
	return &b
}

func False() *bool {
	b := false
	return &b
}