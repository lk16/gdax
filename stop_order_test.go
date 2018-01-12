package gdax

import (
	"testing"
	"encoding/json"
	"context"
)

var stopOrderMarshalCases = []struct {
	input    StopOrderRequest
	expected string
}{
	{
		input: StopOrderRequest{
			Side:      Sell,
			ProductId: "LTC-USD",
			Price:     RequireDecimalFromString("1"),
			Size:      RequireDecimalFromString("2"),
		},
		expected: `{"type":"limit","side":"buy","product_id":"BTC-USD","price":"1","size":"2"}`,
	},
	{
		input: StopOrderRequest{
			Side:                Sell,
			ProductId:           "LTC-USD",
			Price:               RequireDecimalFromString("1"),
			Size:                RequireDecimalFromString("2"),
			ClientOID:           "custom-oid",
			SelfTradePrevention: DecrementAndCancel,
		},
		expected: `{"type":"limit","side":"buy","product_id":"BTC-USD","price":"1","size":"2","client_oid":"custom-oid","stp":"dc","time_in_force":"GTT","cancel_after":"hour","post_only":true}`,
	},
	{
		input: StopOrderRequest{
			Side:                Sell,
			ProductId:           "LTC-USD",
			Price:               RequireDecimalFromString("1"),
			Size:                RequireDecimalFromString("2"),
			ClientOID:           "custom-oid",
			SelfTradePrevention: CancelBoth,
		},
		expected: `{"type":"limit","side":"buy","product_id":"BTC-USD","price":"1","size":"2","client_oid":"custom-oid","stp":"cb","post_only":false}`,
	},
}

func TestStopOrderRequest_MarshalJSON(t *testing.T) {
	for i, c := range stopOrderMarshalCases {
		c.input.Type = Stop
		data, err := json.Marshal(&c.input)
		if err != nil {
			t.Errorf("case %d: could not marshal %s", i, err)
		} else if string(data) != c.expected {
			t.Errorf("case %d: output for did not match expected, was %s", i, string(data))
		}
	}
}

func TestCreateStopOrder_Minimal(t *testing.T) {
	t.Skip("gdax sandbox is down")

	defer testReadWriteClient().CancelAllOrders(context.Background())

	orderRequest := StopOrderRequest{
		Side:      Sell,
		ProductId: "BTC-USD",
		Price:     RequireDecimalFromString("1.00"),
		Size:      RequireDecimalFromString("2.00"),
	}

	orderResponse, err := testReadWriteClient().CreateStopOrder(context.Background(), &orderRequest)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.ID == "" {
		t.Error("order id missing, something was probably incorrect")
	}

	err = compareFields(orderRequest, orderResponse, []string{"Side", "ProductId"})
	if err != nil {
		t.Error(err)
	}
}
