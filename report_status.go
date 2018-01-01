package gdax

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ReportStatus byte

const (
	ReportStatusPending ReportStatus = iota
	ReportStatusCreating
	ReportStatusReady
)

func (d *ReportStatus) MarshalJSON() ([]byte, error) {
	switch *d {
	case ReportStatusPending:
		return json.Marshal("pending")
	case ReportStatusCreating:
		return json.Marshal("creating")
	case ReportStatusReady:
		return json.Marshal("ready")
	}
	return nil, errors.New(fmt.Sprintf("unsupported Type ordinal: %d", *d))
}

func (d *ReportStatus) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	switch s {
	case "pending":
		*d = ReportStatusPending
	case "creating":
		*d = ReportStatusCreating
	case "ready":
		*d = ReportStatusReady
	default:
		err = errors.New(fmt.Sprintf("unsupported Type string: '%s'", s))
	}
	return err
}
