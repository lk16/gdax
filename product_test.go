package gdax

import (
	"context"
	"testing"
	"time"
)

func TestGetProducts(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	products, err := client.GetProducts(context.Background())
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
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	_, err := client.GetBook(context.Background(), "BTC-USD", 1)
	if err != nil {
		t.Error(err)
	}
	_, err = client.GetBook(context.Background(), "BTC-USD", 2)
	if err != nil {
		t.Error(err)
	}
	_, err = client.GetBook(context.Background(), "BTC-USD", 3)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTicker(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	ticker, err := client.GetTicker(context.Background(), "BTC-USD")
	if err != nil {
		t.Error(err)
	}

	if structHasZeroValues(ticker) {
		t.Error("zero value")
	}
}

func TestListTrades(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	var trades []Trade
	cursor := client.ListTrades(context.Background(), "BTC-USD")

	for cursor.HasMore {
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
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	params := GetHistoricRatesParams{
		Start:       time.Now().Add(-24 * 4 * time.Hour),
		End:         time.Now().Add(-24 * 2 * time.Hour),
		Granularity: 1000,
	}

	_, err := client.GetHistoricRates(context.Background(), "BTC-USD", params)
	if err != nil {
		t.Error(err)
	}
}

func TestGetStats(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	stats, err := client.GetStats(context.Background(), "BTC-USD")
	if err != nil {
		t.Error(err)
	}
	if structHasZeroValues(stats) {
		t.Error("zero value")
	}
}
