package gdax

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type ServerTime struct {
	ISO   string  `json:"iso"`
	Epoch float64 `json:"epoch,number"` // decimal seconds since Unix Epoch
}

// GetTime gets the API server time.
func (c *Client) GetTime(ctx context.Context) (ServerTime, error) {
	var serverTime ServerTime

	url := fmt.Sprintf("/time")
	_, err := c.request(ctx, false, "GET", url, nil, &serverTime)
	return serverTime, err
}

var timeLayouts = []string{
	"2006-01-02 15:04:05+00",
	"2006-01-02T15:04:05.999999Z",

	"2006-01-02T15:04:05.999",
	"2006-01-02 15:04:05.999999",
	"2006-01-02T15:04:05Z",
	"2006-01-02 15:04:05.999999+00",
}

func parseTime(s string) (t time.Time, err error) {
	for _, layout := range timeLayouts {
		t, err = time.Parse(layout, s)
		if err == nil {
			break
		}
	}
	return
}

func mustParseTime(s string) time.Time {
	if t, err := parseTime(s); err != nil {
		panic(err)
	} else {
		return t
	}
}

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*t = Time(time.Time{})
		return nil
	}

	dataNoQuotes := strings.Replace(string(data), "\"", "", -1)
	if parsedTime, err := parseTime(dataNoQuotes); err != nil {
		return err
	} else {
		*t = Time(parsedTime)
		return nil
	}
}

// MarshalJSON marshal time back to time.Time for json encoding
func (t Time) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}

func (t *Time) Time() time.Time {
	return time.Time(*t)
}
