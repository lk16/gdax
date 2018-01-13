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

func BoolRef(b bool) *bool {
	return &b
}

func True() *bool {
	return BoolRef(true)
}

func False() *bool {
	return BoolRef(false)
}

func IDRef(id uuid.UUID) *uuid.UUID {
	return &id;
}

func DecimalRef(d decimal.Decimal) *decimal.Decimal {
	return &d
}
