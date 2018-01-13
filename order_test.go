package gdax

import (
	"context"
	"errors"
	//"github.com/google/uuid"
	"testing"
	"github.com/google/uuid"
)

func TestCreateLimitOrders(t *testing.T) {
	t.Skip("gdax sandbox is down")

	order := OrderRequest{
		Price:     RequireDecimalFromString("1.00"),
		Size:      RequireDecimalFromString("1.00"),
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := testReadWriteClient().CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.ID == uuid.Nil {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	err = compareFields(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateMarketOrders(t *testing.T) {
	t.Skip("gdax sandbox is down")

	order := OrderRequest{
		Funds:     "10.00",
		Size:      RequireDecimalFromString("2.00"),
		Side:      "buy",
		Type:      "market",
		ProductId: "BTC-USD",
	}

	savedOrder, err := testReadWriteClient().CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	if savedOrder.ID == uuid.Nil {
		t.Error(errors.New("No create id found"))
	}

	props := []string{"Price", "Size", "Side", "ProductId"}
	err = compareFields(order, savedOrder, props)
	if err != nil {
		t.Error(err)
	}
}

func TestCancelOrder(t *testing.T) {
	t.Skip("gdax sandbox is down")

	var orders []OrderResponse
	cursor := testReadWriteClient().ListOrders(context.Background())
	for cursor.HasMore {
		if err := cursor.NextPage(&orders); err != nil {
			t.Error(err)
		}

		for _, o := range orders {
			if err := testReadWriteClient().CancelOrder(context.Background(), o.ID); err != nil {
				if err.Error() != "OrderRequest already done" {
					t.Error(err)
				}
			}
		}
	}
}

func TestGetOrder(t *testing.T) {
	t.Skip("gdax sandbox is down")

	order := OrderRequest{
		Price:     RequireDecimalFromString("1.00"),
		Size:      RequireDecimalFromString("1.00"),
		Side:      "buy",
		ProductId: "BTC-USD",
	}

	savedOrder, err := testReadWriteClient().CreateOrder(context.Background(), &order)
	if err != nil {
		t.Error(err)
	}

	getOrder, err := testReadWriteClient().GetOrder(context.Background(), savedOrder.ID)
	if err != nil {
		t.Error(err)
	}

	if getOrder.ID != savedOrder.ID {
		t.Error(errors.New("OrderRequest ids do not match"))
	}
}

func TestListOrders(t *testing.T) {
	cursor := testReadOnlyClient().ListOrders(context.Background())
	var orders []OrderRequest

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
		order := OrderRequest{Price: RequireDecimalFromString("1.00"), Size: RequireDecimalFromString("10000.00"), Side: "buy", ProductId: pair}

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
