package gdax

import (
	"context"
	"testing"
)

func TestListFills(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	cursor := client.ListFills(context.Background())
	var fills []Fill

	for cursor.HasMore {
		if err := cursor.NextPage(&fills); err != nil {
			t.Error(err)
		}

		for _, f := range fills {
			if structHasZeroValues(f) {
				t.Error("zero value")
			}
		}
	}
	params := ListFillsParams{
		ProductId: "BTC-USD",
	}
	cursor = client.ListFills(context.Background(), params)
	for cursor.HasMore {
		if err := cursor.NextPage(&fills); err != nil {
			t.Error(err)
		}

		for _, f := range fills {
			if structHasZeroValues(f) {
				t.Error("zero value")
			}
		}
	}
}
