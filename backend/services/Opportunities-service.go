package services

import (
	"time"
)

// OpportunityRequest servis modeli
type OpportunityRequest struct {
	AnalysisInfoId uint      `json:"analysis_info_id" validate:"required"` // Foreign key
	Opportunity    string    `json:"opportunity" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
