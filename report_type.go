package gdax

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ReportType byte

const (
	ReportTypeFills ReportType = iota
	ReportTypeAccount
)

func (d *ReportType) MarshalJSON() ([]byte, error) {
	switch *d {
	case ReportTypeFills:
		return json.Marshal("fills")
	case ReportTypeAccount:
		return json.Marshal("account")
	}
	return nil, errors.New(fmt.Sprintf("unsupported ReportType ordinal: %d", *d))
}

func (d *ReportType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	switch s {
	case "fills":
		*d = ReportTypeFills
		break
	case "account":
		*d = ReportTypeAccount
		break
	default:
		err = errors.New(fmt.Sprintf("unsupported ReportType string: '%s'", s))
	}
	return err
}
