package gdax

import (
	"testing"
	"encoding/json"
	"context"
	"github.com/google/uuid"
)

var limitOrderMarshalCases = []struct {
	input    LimitOrderRequest
	expected string
}{
	{
		input: LimitOrderRequest{
			Side:      Buy,
			ProductId: "BTC-USD",
			Price:     RequireDecimalFromString("1"),
			Size:      RequireDecimalFromString("2"),
		},
		expected: `{"type":"limit","side":"buy","product_id":"BTC-USD","price":"1","size":"2"}`,
	},
	{
		input: LimitOrderRequest{
			Side:                Buy,
			ProductId:           "BTC-USD",
			Price:               RequireDecimalFromString("1"),
			Size:                RequireDecimalFromString("2"),
			ClientOID:           "custom-oid",
			SelfTradePrevention: DecrementAndCancel,
			TimeInForce:         GoodTillTime,
			CancelAfter:         CancelAfterHour,
			PostOnly:            True(),
		},
		expected: `{"type":"limit","side":"buy","product_id":"BTC-USD","price":"1","size":"2","client_oid":"custom-oid","stp":"dc","time_in_force":"GTT","cancel_after":"hour","post_only":true}`,
	},
	{
		input: LimitOrderRequest{
			Side:                Buy,
			ProductId:           "BTC-USD",
			Price:               RequireDecimalFromString("1"),
			Size:                RequireDecimalFromString("2"),
			ClientOID:           "custom-oid",
			SelfTradePrevention: CancelBoth,
			PostOnly:            False(),
		},
		expected: `{"type":"limit","side":"buy","product_id":"BTC-USD","price":"1","size":"2","client_oid":"custom-oid","stp":"cb","post_only":false}`,
	},
}

func TestLimitOrderRequest_MarshalJSON(t *testing.T) {
	for i, c := range limitOrderMarshalCases {
		c.input.Type = Limit
		data, err := json.Marshal(&c.input)
		if err != nil {
			t.Errorf("case %d: could not marshal %s", i, err)
		} else if string(data) != c.expected {
			t.Errorf("case %d: output for did not match expected, was %s", i, string(data))
		}
	}
}

func TestLimitOrderResponse_UnmarshalJSON(t *testing.T) {
	response := `{
		"created_at": "2018-01-12T15:43:53.000000Z",
		"executed_value": "0.0000000000000000",
		"fill_fees": "0.0000000000000000",
		"filled_size": "0.00000000",
		"id": "9e2412cf-1952-49cb-9173-4ea48a5508fb",
		"post_only": false,
		"price": "1.00000000",
		"product_id": "BTC-USD",
		"settled": false,
		"side": "buy",
		"size": "2.00000000",
		"status": "pending",
		"stp": "dc",
		"time_in_force": "GTC",
		"type": "limit"
	}`
	expected := LimitOrderResponse{
		CreatedAt:           Time(timeMust(parseTime("2018-01-12T15:43:53.000000Z"))),
		ExecutedValue:       RequireDecimalFromString("0"),
		FillFees:            RequireDecimalFromString("0"),
		FilledSize:          RequireDecimalFromString("0"),
		ID:                  uuid.Must(uuid.Parse("9e2412cf-1952-49cb-9173-4ea48a5508fb")),
		PostOnly:            false,
		Price:               RequireDecimalFromString("1"),
		ProductId:           "BTC-USD",
		Settled:             false,
		Side:                Buy,
		Size:                RequireDecimalFromString("2"),
		Status:              OrderStatusPending,
		SelfTradePrevention: DecrementAndCancel,
		TimeInForce:         GoodTillCanceled,
		Type:                Limit,
	}
	var actual LimitOrderResponse
	if err := json.Unmarshal([]byte(response), &actual); err != nil {
		t.Errorf("could not unmarshal order response json: %s", err)
	} else if err := compareAllFields(expected, actual); err != nil {
		t.Error(err)
	}
}

func TestCreateLimitOrder_Minimal(t *testing.T) {
	t.Skip("gdax sandbox is down")

	defer testReadWriteClient().CancelAllOrders(context.Background())

	orderRequest := LimitOrderRequest{
		Side:      Buy,
		ProductId: "BTC-USD",
		Price:     RequireDecimalFromString("1.00"),
		Size:      RequireDecimalFromString("2.00"),
	}

	orderResponse, err := testReadWriteClient().CreateLimitOrder(context.Background(), &orderRequest)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.ID == uuid.Nil {
		t.Error("order id missing, something was probably incorrect")
	}

	if err = compareFields(orderRequest, orderResponse, []string{"Side", "ProductId"}); err != nil {
		t.Error(err)
	}
}
