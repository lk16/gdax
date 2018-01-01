package report

import (
	"testing"
)

func TestReportTypeMarshalJson (t *testing.T) {
	reportType := TypeFills
	bytes, err := (&reportType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"fills"` {
		t.Errorf(`was expecting TypeFills to marshal to "fills", was %s`, string(bytes))
	}

	reportType = TypeAccount
	bytes, err = (&reportType).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"account"` {
		t.Errorf(`was expecting TypeFills to marshal to "account", was %s`, string(bytes))
	}
}

func TestReportTypeMarshalJsonErr (t *testing.T) {
	reportType := Type(4)
	_, err := (&reportType).MarshalJSON()
	if err == nil {
		t.Errorf("was expecting a marshalling error")
	}
}

func TestReportTypeUnmarshalJson (t *testing.T) {
	var reportType Type

	err := (&reportType).UnmarshalJSON([]byte(`"fills"`))
	if err != nil {
		t.Error(err)
	}
	if reportType != TypeFills {
		t.Errorf("was expecting TypeFills")
	}

	err = (&reportType).UnmarshalJSON([]byte(`"account"`))
	if err != nil {
		t.Error(err)
	}
	if reportType != TypeAccount {
		t.Errorf("was expecting TypeAccount")
	}
}

func TestReportTypeUnmarshalJsonErr (t *testing.T) {
	var reportType Type
	err := (&reportType).UnmarshalJSON([]byte(`"not__valid_report_type"`))
	if err == nil {
		t.Error("was expecting and unmarshalling error")
	}
}