package gdax

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	serverTime, err := testClient().GetTime(context.Background())
	if err != nil {
		t.Error(err)
	}

	if structHasZeroValues(serverTime) {
		t.Error("zero value")
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	c := Time{}
	now := time.Now()

	jsonData, err := json.Marshal(now.Format("2006-01-02 15:04:05+00"))
	if err != nil {
		t.Error(err)
	}

	if err = c.UnmarshalJSON(jsonData); err != nil {
		t.Error(err)
	}

	if now.Equal(c.Time()) {
		t.Errorf("unmarshaled time (%s) does not equal original time (%s)", now, c.Time())
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	c := Time{}
	tt := time.Date(9999, 4, 12, 23, 20, 50, 0, time.UTC)
	expected := "\"9999-04-12T23:20:50Z\""

	jsonData, err := json.Marshal(tt.Format("2006-01-02 15:04:05+00"))
	if err != nil {
		t.Error(err)
	}

	if err = c.UnmarshalJSON(jsonData); err != nil {
		t.Error(err)
	}

	jsonData, err = json.Marshal(c)
	if err != nil {
		t.Error(err)
	}

	if string(jsonData) != expected {
		t.Errorf("marshaled time (%s) does not equal original time (%s)", string(jsonData), expected)
	}
}
