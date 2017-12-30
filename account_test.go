package gdax

import (
	"context"
	"testing"
)

func TestGetAccounts(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	accounts, err := client.GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	// Check for decoding issues
	for _, a := range accounts {
		if structHasZeroValues(a) {
			t.Error("zero value")
		}
	}
}

func TestGetAccount(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	accounts, err := client.GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, a := range accounts {
		account, err := client.GetAccount(context.Background(), a.Id)
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
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	var ledger []LedgerEntry
	accounts, err := client.GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, a := range accounts {
		cursor := client.ListAccountLedger(context.Background(), a.Id)
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
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	var holds []Hold
	accounts, err := client.GetAccounts(context.Background())
	if err != nil {
		t.Error(err)
	}

	for _, a := range accounts {
		cursor := client.ListHolds(context.Background(), a.Id)
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
