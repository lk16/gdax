package gdax

import (
	"context"
	"testing"
	"time"
)

func TestCreateReportAndStatus(t *testing.T) {
	client := testClient()
	if !client.hasCredentials {
		t.Skip("credentials are required to test")
		return
	}

	newReport := Report{
		Type:      "fill",
		StartDate: time.Now().Add(-24 * 4 * time.Hour),
		EndDate:   time.Now().Add(-24 * 2 * time.Hour),
	}

	report, err := client.CreateReport(context.Background(), &newReport)
	if err != nil {
		t.Error(err)
	}

	currentReport, err := client.GetReportStatus(context.Background(), report.Id)
	if err != nil {
		t.Error(err)
	}
	if structHasZeroValues(currentReport) {
		t.Error("zero value")
	}
}
