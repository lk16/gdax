package gdax

import (
	"github.com/shopspring/decimal"
	"github.com/google/uuid"
)

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

func IDRef(id uuid.UUID) *uuid.UUID {
	return &id;
}