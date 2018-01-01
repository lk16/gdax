package gdax

import (
	"testing"
)

func TestStatusMarshalJson(t *testing.T) {
	status := ReportStatusPending
	bytes, err := (&status).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"pending"` {
		t.Errorf(`was expecting TypeFills to marshal to "fills", was %s`, string(bytes))
	}

	status = ReportStatusCreating
	bytes, err = (&status).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"creating"` {
		t.Errorf(`was expecting TypeFills to marshal to "account", was %s`, string(bytes))
	}

	status = ReportStatusReady
	bytes, err = (&status).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"ready"` {
		t.Errorf(`was expecting TypeFills to marshal to "account", was %s`, string(bytes))
	}
}

func TestStatusMarshalJsonErr(t *testing.T) {
	status := ReportStatus(255)
	_, err := (&status).MarshalJSON()
	if err == nil {
		t.Errorf("was expecting a marshalling error")
	}
}

func TestStatusUnmarshalJson(t *testing.T) {
	var status ReportStatus

	err := (&status).UnmarshalJSON([]byte(`"pending"`))
	if err != nil {
		t.Error(err)
	}
	if status != ReportStatusPending {
		t.Errorf("was expecting ReportStatusPending")
	}

	err = (&status).UnmarshalJSON([]byte(`"creating"`))
	if err != nil {
		t.Error(err)
	}
	if status != ReportStatusCreating {
		t.Errorf("was expecting ReportStatusCreating")
	}

	err = (&status).UnmarshalJSON([]byte(`"ready"`))
	if err != nil {
		t.Error(err)
	}
	if status != ReportStatusReady {
		t.Errorf("was expecting ReportStatusReady")
	}
}

func TestStatusUnmarshalJsonErr(t *testing.T) {
	var status ReportStatus
	err := (&status).UnmarshalJSON([]byte(`"invalid_status"`))
	if err == nil {
		t.Error("was expecting and unmarshalling error")
	}
}
