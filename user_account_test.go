package gdax

import (
	"context"
	"testing"
)

func TestGetTrailingVolume(t *testing.T) {
	t.Skip("Still debugging")

	volume, err := testReadOnlyClient().GetTrailingVolume(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(volume) == 0 {
		t.Error("no volume were found")
	}

	// Check for decoding issues
	for _, a := range volume {
		if structHasZeroValues(a) {
			t.Error("zero value")
		}
	}
}
