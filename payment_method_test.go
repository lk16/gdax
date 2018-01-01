package gdax

import (
	"context"
	"testing"
)

func TestGetPaymentMethods(t *testing.T) {
	paymentMethods, err := testReadOnlyClient().GetPaymentMethods(context.Background())
	if err != nil {
		t.Error(err)
	}

	if len(paymentMethods) == 0 {
		t.Error("no accounts were found")
	}

	// Check for decoding issues
	for _, a := range paymentMethods {
		if structHasZeroValues(a) {
			t.Error("zero value")
		}
	}
}
