package gdax

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LedgerEntryType byte

const (
	LedgerEntryTypeTransfer LedgerEntryType = iota
	LedgerEntryTypeMatch
	LedgerEntryTypeFee
	LedgerEntryTypeRebate
)

func (d *LedgerEntryType) MarshalJSON() ([]byte, error) {
	switch *d {
	case LedgerEntryTypeTransfer:
		return json.Marshal("transfer")
	case LedgerEntryTypeMatch:
		return json.Marshal("match")
	case LedgerEntryTypeFee:
		return json.Marshal("fee")
	case LedgerEntryTypeRebate:
		return json.Marshal("rebate")
	}
	return nil, errors.New(fmt.Sprintf("unsupported Type ordinal: %d", *d))
}

func (d *LedgerEntryType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	switch s {
	case "transfer":
		*d = LedgerEntryTypeTransfer
	case "match":
		*d = LedgerEntryTypeMatch
	case "fee":
		*d = LedgerEntryTypeFee
	case "rebate":
		*d = LedgerEntryTypeRebate
	default:
		return errors.New(fmt.Sprintf("unsupported Type string: '%s'", s))
	}
	return nil
}
