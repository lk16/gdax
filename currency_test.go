package gdax

import (
	"context"
	"testing"
)

func TestGetCurrencies(t *testing.T) {
	currencies, err := testPublicClient().GetCurrencies(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, c := range currencies {
		if structHasZeroValues(c) {
			t.Error("zero value")
		}
	}
}
