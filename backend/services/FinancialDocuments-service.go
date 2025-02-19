package services

import (
	"time"
)

// FinancialDocumentsRequest service model
type FinancialDocumentsRequest struct {
	CampaignId          uint      `json:"campaign_id" validate:"required"` // Foreign key
	Subject             string    `json:"subject" validate:"required"`     // Subject title
	DocumentKey         string    `json:"document_key" validate:"required"`// Document key
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
