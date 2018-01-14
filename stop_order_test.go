package gdax

import (
	"testing"
	"encoding/json"
	"context"
	"github.com/google/uuid"
)

func TestStopOrderRequest_MarshalJSON(t *testing.T) {
	cases := []struct {
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

	for i, c := range cases {
		c.input.Type = Stop
		data, err := json.Marshal(&c.input)
		if err != nil {
			t.Errorf("case %d: could not marshal %s", i, err)
		} else if string(data) != c.expected {
			t.Errorf("case %d: output for did not match expected, was %s", i, string(data))
		}
	}
}

func TestStopOrderResponse_UnmarshalJSON(t *testing.T) {
	response := `{
	  "id": "37e9832d-be72-4dc0-8858-22ceae243394",
	  "price": "240.00000000",
	  "size": "0.01000000",
	  "product_id": "LTC-USD",
	  "side": "sell",
	  "stp": "dc",
	  "type": "stop",
	  "post_only": false,
	  "created_at": "2018-01-13T23:42:13.486093Z",
	  "fill_fees": "0.0000000000000000",
	  "filled_size": "0.00000000",
	  "executed_value": "0.0000000000000000",
	  "status": "pending",
	  "settled": false
	}`
	expected := StopOrderResponse{
		CreatedAt:           Time(timeMust(parseTime("2018-01-13T23:42:13.486093Z"))),
		ExecutedValue:       RequireDecimalFromString("0"),
		FillFees:            RequireDecimalFromString("0"),
		FilledSize:          RequireDecimalFromString("0"),
		ID:                  uuid.Must(uuid.Parse("37e9832d-be72-4dc0-8858-22ceae243394")),
		PostOnly:            false,
		Price:               RequireDecimalFromString("240.0"),
		ProductId:           "LTC-USD",
		Settled:             false,
		Side:                Sell,
		Size:                DecimalRef(RequireDecimalFromString("0.01")),
		Status:              OrderStatusPending,
		SelfTradePrevention: DecrementAndCancel,
		Type:                Stop,
	}
	var actual StopOrderResponse
	if err := json.Unmarshal([]byte(response), &actual); err != nil {
		t.Errorf("could not unmarshal order response json: %s", err)
	} else if err := compareAllFields(expected, actual); err != nil {
		t.Error(err)
	}
}

var stopOrderTestCompareFields = []string{
	"Type", "Side", "ProductId", "Price",
}

func TestCreateStopOrder_MinimalSell(t *testing.T) {
	t.Skip("gdax sandbox is down")

	orderRequest := StopOrderRequest{
		Side:      Sell,
		ProductId: "LTC-USD",
		Price:     RequireDecimalFromString("240.00"),
		Size:      DecimalRef(RequireDecimalFromString("0.01")),
	}

	orderResponse, err := testReadWriteClient().CreateStopOrder(context.Background(), &orderRequest)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.ID == uuid.Nil {
		t.Error("order id missing, something was probably incorrect")
	} else {
		defer testReadWriteClient().CancelOrder(context.Background(), orderResponse.ID)
	}

	err = compareFields(orderRequest, orderResponse, stopOrderTestCompareFields)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.Size == nil {
		t.Error("responze Size was nil but expected something")
	} else if orderResponse.Size.String() != orderRequest.Size.String() {
		t.Errorf("responze Size (%s) did not match request Size (%s)", orderResponse.Size.String(), orderRequest.Size.String())
	}

	if orderResponse.Funds != nil {
		t.Errorf("wasn't expecting funds but received (%s)", orderResponse.Funds.String())
	}
}

func TestCreateStopOrder_MinimalBuy(t *testing.T) {
	t.Skip("gdax sandbox is down")

	orderRequest := StopOrderRequest{
		Side:      Buy,
		ProductId: "LTC-USD",
		Price:     RequireDecimalFromString("999.00"),
		Funds:     DecimalRef(RequireDecimalFromString("1")),
	}

	orderResponse, err := testReadWriteClient().CreateStopOrder(context.Background(), &orderRequest)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.ID == uuid.Nil {
		t.Error("order id missing, something was probably incorrect")
	} else {
		defer testReadWriteClient().CancelOrder(context.Background(), orderResponse.ID)
	}

	err = compareFields(orderRequest, orderResponse, stopOrderTestCompareFields)
	if err != nil {
		t.Error(err)
	}

	expectedFunds := orderRequest.Funds.Mul(RequireDecimalFromString("0.997"))
	if orderResponse.Funds == nil {
		t.Error("response Funds was nil but expected something")
	} else if orderResponse.Funds.StringFixed(3) != expectedFunds.StringFixed(3) {
		t.Errorf("response Funds (%s) did not match expected Funds (%s)", orderResponse.Funds.StringFixed(3), expectedFunds.StringFixed(3))
	}

	if orderResponse.Size != nil {
		t.Errorf("wasn't expecting size but received (%s)", orderResponse.Size.String())
	}
}

func TestCreateStopOrder_Sell(t *testing.T) {
	t.Skip("gdax sandbox is down")

	orderRequest := StopOrderRequest{
		Side:                Sell,
		ProductId:           "LTC-USD",
		Price:               RequireDecimalFromString("240.00"),
		Size:                DecimalRef(RequireDecimalFromString("0.01")),
		ClientOID:           IDRef(uuid.New()),
		SelfTradePrevention: CancelBoth,
	}

	orderResponse, err := testReadWriteClient().CreateStopOrder(context.Background(), &orderRequest)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.ID == uuid.Nil {
		t.Error("order id missing, something was probably incorrect")
	} else {
		defer testReadWriteClient().CancelOrder(context.Background(), orderResponse.ID)
	}

	err = compareFields(orderRequest, orderResponse, stopOrderTestCompareFields)
	if err != nil {
		t.Error(err)
	}

	if orderResponse.Size == nil {
		t.Error("responze Size was nil but expected something")
	} else if orderResponse.Size.String() != orderRequest.Size.String() {
		t.Errorf("responze Size (%s) did not match request Size (%s)", orderResponse.Size.String(), orderRequest.Size.String())
	}

	if orderResponse.Funds != nil {
		t.Errorf("wasn't expecting funds but received (%s)", orderResponse.Funds.String())
	}
}
