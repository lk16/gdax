package gdax

import (
	"context"
	"testing"
	"time"
)

func TestCreateReportAndStatus(t *testing.T) {
	t.Skip("gdax sandbox is down")

	newReport := Report{
		Type:      "fill",
		StartDate: time.Now().Add(-24 * 4 * time.Hour),
		EndDate:   time.Now().Add(-24 * 2 * time.Hour),
	}

	report, err := testReadWriteClient().CreateReport(context.Background(), &newReport)
	if err != nil {
		t.Error(err)
	}

	currentReport, err := testReadWriteClient().GetReportStatus(context.Background(), report.Id)
	if err != nil {
		t.Error(err)
	}
	if structHasZeroValues(currentReport) {
		t.Error("zero value")
	}
}
