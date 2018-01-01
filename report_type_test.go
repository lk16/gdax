package gdax

import (
	"testing"
)

func TestTypeMarshalJson(t *testing.T) {
	reportType := ReportTypeFills
	bytes, err := (&reportType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"fills"` {
		t.Errorf(`was expecting ReportTypeFills to marshal to "fills", was %s`, string(bytes))
	}

	reportType = ReportTypeAccount
	bytes, err = (&reportType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"account"` {
		t.Errorf(`was expecting ReportTypeAccount to marshal to "account", was %s`, string(bytes))
	}
}

func TestTypeMarshalJsonErr(t *testing.T) {
	reportType := ReportType(255)
	_, err := (&reportType).MarshalJSON()
	if err == nil {
		t.Errorf("was expecting a marshalling error")
	}
}

func TestTypeUnmarshalJson(t *testing.T) {
	var reportType ReportType

	err := (&reportType).UnmarshalJSON([]byte(`"fills"`))
	if err != nil {
		t.Error(err)
	}
	if reportType != ReportTypeFills {
		t.Errorf("was expecting ReportTypeFills")
	}

	err = (&reportType).UnmarshalJSON([]byte(`"account"`))
	if err != nil {
		t.Error(err)
	}
	if reportType != ReportTypeAccount {
		t.Errorf("was expecting ReportTypeAccount")
	}
}

func TestTypeUnmarshalJsonErr(t *testing.T) {
	var reportType ReportType
	err := (&reportType).UnmarshalJSON([]byte(`"invalid_report_type"`))
	if err == nil {
		t.Error("was expecting and unmarshalling error")
	}
}
