package services

import (
	"time"
)

// ConsPlanRequest servis modeli
type ConsPlanRequest struct {
	ConPlanId      uint      `json:"con_plan_id"`
	AnalysisInfoId uint      `json:"analysis_info_id" validate:"required"` // Foreign key
	ConPlan        string    `json:"con_plan" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

