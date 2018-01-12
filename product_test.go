package gdax

import (
	"context"
	"testing"
	"time"
)

func TestGetProducts(t *testing.T) {
	products, err := testPublicClient().GetProducts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, p := range products {
		if structHasZeroValues(p) {
			t.Error("zero value")
		}
	}
}

func TestGetBook(t *testing.T) {
	for level := range []int{1, 2, 3} {
		_, err := testPublicClient().GetBook(context.Background(), "BTC-USD", level)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetTicker(t *testing.T) {
	ticker, err := testPublicClient().GetTicker(context.Background(), "BTC-USD")
	if err != nil {
		t.Error(err)
	}

	if structHasZeroValues(ticker) {
		t.Error("zero value")
	}
}

func TestListTrades(t *testing.T) {
	var trades []Trade
	cursor := testPublicClient().ListTrades(context.Background(), "BTC-USD")

	if cursor.HasMore {
		if err := cursor.NextPage(&trades); err != nil {
			t.Error(err)
		}

		for _, a := range trades {
			if structHasZeroValues(a) {
				t.Error("zero value")
			}
		}
	}
}

func TestGetHistoricRates(t *testing.T) {
	params := GetHistoricRatesParams{
		Start:       time.Now().Add(-24 * 4 * time.Hour),
		End:         time.Now().Add(-24 * 2 * time.Hour),
		Granularity: 3600,
	}

	_, err := testPublicClient().GetHistoricRates(context.Background(), "BTC-USD", params)
	if err != nil {
		t.Error(err)
	}
}

func TestGetStats(t *testing.T) {
	stats, err := testPublicClient().GetStats(context.Background(), "BTC-USD")
	if err != nil {
		t.Error(err)
	}
	if structHasZeroValues(stats) {
		t.Error("zero value")
	}
}
