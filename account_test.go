package gdax

import (
	"context"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	accounts, err := testReadOnlyClient().GetAccounts(context.Background())
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

func TestGetAccount(t *testing.T) {
	accounts, err := testReadOnlyClient().GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, a := range accounts {
		account, err := testReadOnlyClient().GetAccount(context.Background(), a.Id)
		if err != nil {
			t.Error(err)
		}

		// Check for decoding issues
		if structHasZeroValues(account) {
			t.Error("zero value")
		}
	}
}
func TestListAccountLedger(t *testing.T) {
	var ledger []LedgerEntry
	accounts, err := testReadOnlyClient().GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, a := range accounts {
		cursor := testReadOnlyClient().ListAccountLedger(context.Background(), a.Id)
		for cursor.HasMore {
			if err := cursor.NextPage(&ledger); err != nil {
				t.Error(err)
			}

			for _, l := range ledger {
				// Check for decoding issues
				if structHasZeroValues(l) {
					t.Error("zero value")
				}
			}
		}
	}
}

func TestListHolds(t *testing.T) {
	var holds []Hold
	accounts, err := testReadOnlyClient().GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, a := range accounts {
		cursor := testReadOnlyClient().ListHolds(context.Background(), a.Id)
		for cursor.HasMore {
			if err := cursor.NextPage(&holds); err != nil {
				t.Error(err)
			}

			for _, h := range holds {
				// Check for decoding issues
				if structHasZeroValues(h) {
					t.Error("zero value")
				}
			}
		}
	}
}
