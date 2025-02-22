package services

import (
	"time"
)

// ExplanationsRequest service model
type ExplanationsRequest struct {
	CampaignId    uint      `json:"campaign_id" validate:"required"` // Foreign key
	Explanation   string    `json:"explanation" validate:"required"` // Explanation text
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
