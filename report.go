package gdax

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type ReportType string

const (
	ReportTypeFills   ReportType = "fills"
	ReportTypeAccount ReportType = "account"
)

type ReportFormat string

const (
	ReportFormatPDF ReportFormat = "pdf"
	ReportFormatCSV ReportFormat = "csv"
)

type CreateReport struct {
	Type      ReportType   `json:"type"`
	StartDate time.Time    `json:"start_date,string"`
	EndDate   time.Time    `json:"end_date,string"`
	ProductID string       `json:"product_id,string,omitempty"`
	AccountID *uuid.UUID   `json:"account_id,string,omitempty"`
	Format    ReportFormat `json:"format,omitempty"`
	Email     string       `json:"email,string,omitempty"`
}

type ReportParams struct {
	StartDate time.Time `json:"start_date,string"`
	EndDate   time.Time `json:"end_date,string"`
}

type ReportStatus string

const (
	ReportStatusPending  ReportStatus = "pending"
	ReportStatusCreating ReportStatus = "creating"
	ReportStatusReady    ReportStatus = "ready"
)

type Report struct {
	ID          uuid.UUID    `json:"id,string,omitempty"`
	Type        ReportType   `json:"type"`
	Status      ReportStatus `json:"status"`
	CreatedAt   Time         `json:"created_at,string"`
	CompletedAt Time         `json:"completed_at,string,"`
	ExpiresAt   Time         `json:"expires_at,string"`
	FileURL     string       `json:"file_url"`
	Params      ReportParams `json:"params"`
}

func (c *Client) CreateReport(ctx context.Context, newReport *CreateReport) (Report, error) {
	var report Report

	url := fmt.Sprintf("/reports")
	_, err := c.request(ctx, true, "POST", url, newReport, &report)

	return report, err
}

func (c *Client) GetReportStatus(ctx context.Context, id uuid.UUID) (Report, error) {
	report := Report{}

	url := fmt.Sprintf("/reports/%s", id)
	_, err := c.request(ctx, true, "GET", url, nil, &report)

	return report, err
}
