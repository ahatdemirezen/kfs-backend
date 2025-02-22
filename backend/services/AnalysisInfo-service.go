package services

import (
	"time"
)

// AnalysisInfoRequest servis modeli
type AnalysisInfoRequest struct {
	CampaignId     uint      `json:"campaign_id" validate:"required"` // Foreign key
	SwotKey        string    `json:"swot_key" validate:"required"`
	BusinessKey    string    `json:"business_key" validate:"required"`
	InvestorKey    string    `json:"investor_key" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

