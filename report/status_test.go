package report

import (
	"testing"
)

func TestStatusMarshalJson (t *testing.T) {
	status := StatusPending
	bytes, err := (&status).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"pending"` {
		t.Errorf(`was expecting TypeFills to marshal to "fills", was %s`, string(bytes))
	}

	status = StatusCreating
	bytes, err = (&status).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"creating"` {
		t.Errorf(`was expecting TypeFills to marshal to "account", was %s`, string(bytes))
	}

	status = StatusReady
	bytes, err = (&status).MarshalJSON()
	if err != nil {
		t.Error(err)
	}
	if string(bytes) != `"ready"` {
		t.Errorf(`was expecting TypeFills to marshal to "account", was %s`, string(bytes))
	}
}

func TestStatusMarshalJsonErr (t *testing.T) {
	status := Status(255)
	_, err := (&status).MarshalJSON()
	if err == nil {
		t.Errorf("was expecting a marshalling error")
	}
}

func TestStatusUnmarshalJson (t *testing.T) {
	var status Status

	err := (&status).UnmarshalJSON([]byte(`"pending"`))
	if err != nil {
		t.Error(err)
	}
	if status != StatusPending {
		t.Errorf("was expecting StatusPending")
	}

	err = (&status).UnmarshalJSON([]byte(`"creating"`))
	if err != nil {
		t.Error(err)
	}
	if status != StatusCreating {
		t.Errorf("was expecting StatusCreating")
	}

	err = (&status).UnmarshalJSON([]byte(`"ready"`))
	if err != nil {
		t.Error(err)
	}
	if status != StatusReady {
		t.Errorf("was expecting StatusReady")
	}
}

func TestStatusUnmarshalJsonErr (t *testing.T) {
	var status Status
	err := (&status).UnmarshalJSON([]byte(`"invalid_status"`))
	if err == nil {
		t.Error("was expecting and unmarshalling error")
	}
}