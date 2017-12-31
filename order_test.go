package gdax

import (
	"context"
	"errors"
	"testing"
)

func TestCreateLimitOrders(t *testing.T) {
	t.Skip("gdax sandbox is down")

	order := Order{
		Price:     requireDecimalFromString("1.00"),
		Size:      requireDecimalFromString("1.00"),
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := testReadWriteClient().CreateOrder(context.Background(), &order)
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
	t.Skip("gdax sandbox is down")

	order := Order{
		Funds:     requireDecimalFromString("10.00"),
		Size:      requireDecimalFromString("2.00"),
		Side:      "buy",
		Type:      "market",
		ProductId: "BTC-USD",
	}

	savedOrder, err := testReadWriteClient().CreateOrder(context.Background(), &order)
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
	t.Skip("gdax sandbox is down")

	var orders []Order
	cursor := testReadWriteClient().ListOrders(context.Background())
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if err := testReadWriteClient().CancelOrder(context.Background(), o.Id); err != nil {
				if err.Error() != "Order already done" {
					t.Error(err)
				}
			}
		}
	}
}

func TestGetOrder(t *testing.T) {
	t.Skip("gdax sandbox is down")

	order := Order{
		Price:     requireDecimalFromString("1.00"),
		Size:      requireDecimalFromString("1.00"),
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := testReadWriteClient().CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	getOrder, err := testReadWriteClient().GetOrder(context.Background(), savedOrder.Id)
	if err != nil {
		t.Error(err)
	}

	if getOrder.Id != savedOrder.Id {
		t.Error(errors.New("Order ids do not match"))
	}
}

func TestListOrders(t *testing.T) {
	cursor := testReadOnlyClient().ListOrders(context.Background())
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

	cursor = testReadOnlyClient().ListOrders(context.Background(), ListOrdersParams{Status: "open", ProductId: "LTC-EUR"})
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
	t.Skip("gdax sandbox is down")

	for _, pair := range []string{"BTC-USD", "ETH-USD", "LTC-USD"} {
		order := Order{Price: requireDecimalFromString("1.00"), Size: requireDecimalFromString("10000.00"), Side: "buy", ProductId: pair}

		if _, err := testReadWriteClient().CreateOrder(context.Background(), &order); err != nil {
			t.Error(err)
		}
	}

	orderIDs, err := testReadWriteClient().CancelAllOrders(context.Background(), CancelAllOrdersParams{ProductId: "LTC-USD"})
	if err != nil {
		t.Error(err)
	}

	if len(orderIDs) != 1 {
		t.Error("did not cancel single order")
	}
}
