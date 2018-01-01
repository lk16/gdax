package report

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Type byte

const (
	TypeFills   Type = iota
	TypeAccount
)

func (d *Type) MarshalJSON() ([]byte, error) {
	switch *d {
	case TypeFills:
		return json.Marshal("fills")
	case TypeAccount:
		return json.Marshal("account")
	}
	return nil, errors.New(fmt.Sprintf("unsupported Type ordinal: %d", *d))
}

func (d *Type) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}

	switch s {
	case "fills":
		*d = TypeFills
		break
	case "account":
		*d = TypeAccount
		break
	default:
		err = errors.New(fmt.Sprintf("unsupported Type string: '%s'", s))
	}
	return err
}
