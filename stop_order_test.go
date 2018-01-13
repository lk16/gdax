package gdax

import (
	"testing"
	"encoding/json"
	"context"
	"github.com/google/uuid"
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
			Size:      DecimalRef(RequireDecimalFromString("2")),
		},
		expected: `{"type":"stop","side":"sell","product_id":"LTC-USD","price":"1","size":"2"}`,
	},
	{
		input: StopOrderRequest{
			Side:                Sell,
			ProductId:           "LTC-USD",
			Price:               RequireDecimalFromString("1"),
			Size:                DecimalRef(RequireDecimalFromString("2")),
			ClientOID:           IDRef(uuid.Must(uuid.Parse("c35bb839-ac0a-4ae0-aeee-c3b11a8f20c5"))),
			SelfTradePrevention: DecrementAndCancel,
		},
		expected: `{"type":"stop","side":"sell","product_id":"LTC-USD","price":"1","size":"2","client_oid":"c35bb839-ac0a-4ae0-aeee-c3b11a8f20c5","stp":"dc"}`,
	},
	{
		input: StopOrderRequest{
			Side:                Sell,
			ProductId:           "LTC-USD",
			Price:               RequireDecimalFromString("1"),
			Funds:               DecimalRef(RequireDecimalFromString("2")),
			ClientOID:           IDRef(uuid.Must(uuid.Parse("c35bb839-ac0a-4ae0-aeee-c3b11a8f20c5"))),
			SelfTradePrevention: CancelBoth,
		},
		expected: `{"type":"stop","side":"sell","product_id":"LTC-USD","price":"1","funds":"2","client_oid":"c35bb839-ac0a-4ae0-aeee-c3b11a8f20c5","stp":"cb"}`,
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
		Size:      DecimalRef(RequireDecimalFromString("2.00")),
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
