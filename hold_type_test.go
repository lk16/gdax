package gdax

import (
	"testing"
)

func TestHoldTypeMarshalJson(t *testing.T) {
	holdType := HoldTypeOrder
	bytes, err := (&holdType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"order"` {
		t.Errorf(`was LedgerEntryTypeTransfer HoldTypeOrder to marshal to "order", was %s`, string(bytes))
	}

	holdType = HoldTypeTransfer
	bytes, err = (&holdType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"transfer"` {
		t.Errorf(`was expecting HoldTypeTransfer to marshal to "transfer", was %s`, string(bytes))
	}
}

func TestHoldTypeMarshalJsonErr(t *testing.T) {
	holdType := HoldType(255)
	_, err := (&holdType).MarshalJSON()
	if err == nil {
		t.Errorf("was expecting a marshalling error")
	}
}

func TestHoldTypeUnmarshalJson(t *testing.T) {
	var holdType HoldType

	err := (&holdType).UnmarshalJSON([]byte(`"order"`))
	if err != nil {
		t.Error(err)
	}
	if holdType != HoldTypeOrder {
		t.Errorf("was expecting HoldTypeOrder")
	}

	err = (&holdType).UnmarshalJSON([]byte(`"transfer"`))
	if err != nil {
		t.Error(err)
	}
	if holdType != HoldTypeTransfer {
		t.Errorf("was expecting HoldTypeTransfer")
	}
}

func TestHoldTypeUnmarshalJsonErr(t *testing.T) {
	var holdType HoldType
	err := (&holdType).UnmarshalJSON([]byte(`"invalid_hold_type"`))
	if err == nil {
		t.Error("was expecting and unmarshalling error")
	}
}
