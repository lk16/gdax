package gdax

import (
	"context"
	"testing"
)

func TestGetCoinbaseAccounts(t *testing.T) {
	accounts, err := testReadOnlyClient().GetCoinbaseAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(accounts) == 0 {
		t.Error("no accounts were found")
	}

	// Check for decoding issues
	for _, a := range accounts {
		if structHasZeroValues(a) {
			t.Error("zero value")
		}
	}
}
