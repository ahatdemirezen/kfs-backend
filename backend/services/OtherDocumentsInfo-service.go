package services

import (
	"time"
)

// OtherDocumentsInfoRequest service model
type OtherDocumentsInfoRequest struct {
	CampaignId  uint      `json:"campaign_id" validate:"required"` // Foreign key
	DocumentKey string    `json:"document_key" validate:"required"` // Document key
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}


