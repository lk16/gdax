package gdax

import (
	"context"
	"encoding/json"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	serverTime, err := testPublicClient().GetTime(context.Background())
	if err != nil {
		t.Error(err)
	}

	if structHasZeroValues(serverTime) {
		t.Error("zero value")
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {

	cases := []struct {
		input    string
		expected time.Time
	}{
		{
			input:"2018-01-13T17:20:10.601",
			expected:time.Date(2018, 1, 13, 17, 20, 10, 601000000, time.UTC),
		},
		{
			input:"2006-01-02 15:04:05+00",
			expected:time.Date(2006, 1, 2, 15, 4, 5, 0, time.UTC),
		},
	}

	for _, testCase := range cases {
		var c Time
		if err := c.UnmarshalJSON([]byte(testCase.input)); err != nil {
			t.Error(err)
		} else if !testCase.expected.Equal(c.Time()) {
			t.Errorf("unmarshaled time (%s) does not equal expected time (%s)", c.Time(), testCase.expected)
		}
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
