package gdax

import (
	"testing"
)

func TestLedgerEntryTypeMarshalJson(t *testing.T) {
	ledgerEntryType := LedgerEntryTypeTransfer
	bytes, err := (&ledgerEntryType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"transfer"` {
		t.Errorf(`was LedgerEntryTypeTransfer ReportTypeFills to marshal to "transfer", was %s`, string(bytes))
	}

	ledgerEntryType = LedgerEntryTypeMatch
	bytes, err = (&ledgerEntryType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"match"` {
		t.Errorf(`was expecting LedgerEntryTypeMatch to marshal to "match", was %s`, string(bytes))
	}

	ledgerEntryType = LedgerEntryTypeFee
	bytes, err = (&ledgerEntryType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"fee"` {
		t.Errorf(`was expecting LedgerEntryTypeMatch to marshal to "fee", was %s`, string(bytes))
	}

	ledgerEntryType = LedgerEntryTypeRebate
	bytes, err = (&ledgerEntryType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"rebate"` {
		t.Errorf(`was expecting LedgerEntryTypeMatch to marshal to "rebate", was %s`, string(bytes))
	}
}

func TestLedgerEntryTypeMarshalJsonErr(t *testing.T) {
	ledgerEntryType := LedgerEntryType(255)
	_, err := (&ledgerEntryType).MarshalJSON()
	if err == nil {
		t.Errorf("was expecting a marshalling error")
	}
}

func TestLedgerEntryTypeUnmarshalJson(t *testing.T) {
	var ledgerEntryType LedgerEntryType

	err := (&ledgerEntryType).UnmarshalJSON([]byte(`"transfer"`))
	if err != nil {
		t.Error(err)
	}
	if ledgerEntryType != LedgerEntryTypeTransfer {
		t.Errorf("was expecting LedgerEntryTypeTransfer")
	}

	err = (&ledgerEntryType).UnmarshalJSON([]byte(`"match"`))
	if err != nil {
		t.Error(err)
	}
	if ledgerEntryType != LedgerEntryTypeMatch {
		t.Errorf("was expecting LedgerEntryTypeMatch")
	}

	err = (&ledgerEntryType).UnmarshalJSON([]byte(`"fee"`))
	if err != nil {
		t.Error(err)
	}
	if ledgerEntryType != LedgerEntryTypeFee {
		t.Errorf("was expecting LedgerEntryTypeFee")
	}

	err = (&ledgerEntryType).UnmarshalJSON([]byte(`"rebate"`))
	if err != nil {
		t.Error(err)
	}
	if ledgerEntryType != LedgerEntryTypeRebate {
		t.Errorf("was expecting LedgerEntryTypeRebate")
	}
}

func TestTypeUnmarshalJsonErr(t *testing.T) {
	var reportType ReportType
	err := (&reportType).UnmarshalJSON([]byte(`"invalid_ledger_entry_type"`))
	if err == nil {
		t.Error("was expecting and unmarshalling error")
	}
}
