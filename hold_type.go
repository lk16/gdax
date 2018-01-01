package gdax

import (
	"encoding/json"
	"errors"
	"fmt"
)

type HoldType byte

const (
	HoldTypeOrder HoldType = iota
	HoldTypeTransfer
)

func (d *HoldType) MarshalJSON() ([]byte, error) {
	switch *d {
	case HoldTypeOrder:
		return json.Marshal("order")
	case HoldTypeTransfer:
		return json.Marshal("transfer")
	}
	return nil, errors.New(fmt.Sprintf("unsupported HoldType ordinal: %d", *d))
}

func (d *HoldType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case "order":
		*d = HoldTypeOrder
	case "transfer":
		*d = HoldTypeTransfer
	default:
		return errors.New(fmt.Sprintf("unsupported HoldType string: '%s'", s))
	}
	return nil
}

