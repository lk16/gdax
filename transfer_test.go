package gdax

import (
	"context"
	"os"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
	t.Skip("gdax sandbox is down")

	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	transfer := Transfer{
		Type:              "deposit",
		Amount:            requireDecimalFromString("1.00"),
		CoinbaseAccountId: os.Getenv("TEST_COINBASE_ACCOUNT_ID"),
	}
	if transfer.CoinbaseAccountId == "" {
		t.Skip("skipping, TEST_COINBASE_ACCOUNT_ID must be specified")
		return
	}

	depositTransfer, err := client.CreateTransfer(context.Background(), &transfer)
	if err != nil {
		t.Error(err)
		return
	}

	if depositTransfer.CoinbaseAccountId != transfer.CoinbaseAccountId {
		t.Errorf("CoinbaseAccountId did not match the one sent")
	}
	if depositTransfer.Amount != transfer.Amount {
		t.Errorf("Amount did not match the one sent")
	}
	if depositTransfer.Type != transfer.Type {
		t.Errorf("CoinbaseAccountId did not match the one sent")
	}
}
