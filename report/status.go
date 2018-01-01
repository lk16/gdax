package report

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Status byte

const (
	StatusPending  Status = iota
	StatusCreating
	StatusReady
)

func (d *Status) MarshalJSON() ([]byte, error) {
	switch *d {
	case StatusPending:
		return json.Marshal("pending")
	case StatusCreating:
		return json.Marshal("creating")
	case StatusReady:
		return json.Marshal("ready")
	}
	return nil, errors.New(fmt.Sprintf("unsupported Type ordinal: %d", *d))
}

func (d *Status) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	switch s {
	case "pending":
		*d = StatusPending
	case "creating":
		*d = StatusCreating
	case "ready":
		*d = StatusReady
	default:
		err = errors.New(fmt.Sprintf("unsupported Type string: '%s'", s))
	}
	return err
}
