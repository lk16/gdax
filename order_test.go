package gdax

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"testing"
)

func TestCreateLimitOrders(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	order := Order{
		Price:     decimal.RequireFromString("1.00"),
		Size:      decimal.RequireFromString("1.00"),
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.Id == "" {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	_, err = compareProperties(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMarketOrders(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	order := Order{
		Funds:     decimal.RequireFromString("10.00"),
		Size:      decimal.RequireFromString("2.00"),
		Side:      "buy",
		Type:      "market",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.Id == "" {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	_, err = compareProperties(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCancelOrder(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	var orders []Order
	cursor := client.ListOrders(context.Background())
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if err := client.CancelOrder(context.Background(), o.Id); err != nil {
				if err.Error() != "Order already done" {
					t.Error(err)
				}
			}
		}
	}
}

func TestGetOrder(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	order := Order{
		Price:     decimal.RequireFromString("1.00"),
		Size:      decimal.RequireFromString("1.00"),
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := client.CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	getOrder, err := client.GetOrder(context.Background(), savedOrder.Id)
	if err != nil {
		t.Error(err)
	}

	if getOrder.Id != savedOrder.Id {
		t.Error(errors.New("Order ids do not match"))
	}
}

func TestListOrders(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	cursor := client.ListOrders(context.Background())
	var orders []Order

	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if structHasZeroValues(o) {
				t.Error("zero value")
			}
		}
	}

	cursor = client.ListOrders(context.Background(), ListOrdersParams{Status: "open", ProductId: "LTC-EUR"})
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if structHasZeroValues(o) {
				t.Error("zero value")
			}
		}
	}
}

func TestCancelAllOrders(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	for _, pair := range []string{"BTC-USD", "ETH-USD", "LTC-USD"} {
		order := Order{Price: decimal.RequireFromString("1.00"), Size: decimal.RequireFromString("10000.00"), Side: "buy", ProductId: pair}

		if _, err := client.CreateOrder(context.Background(), &order); err != nil {
			t.Error(err)
		}
	}

	orderIDs, err := client.CancelAllOrders(context.Background(), CancelAllOrdersParams{ProductId: "LTC-USD"})
	if err != nil {
		t.Error(err)
	}

	if len(orderIDs) != 1 {
		t.Error("did not cancel single order")
	}
}
