package services

import (
	"time"
)

// ThreatPlanRequest servis modeli
type ThreatPlanRequest struct {
	AnalysisInfoId uint      `json:"analysis_info_id" validate:"required"` // Foreign key
	ThreatPlan     string    `json:"threat_plan" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}


