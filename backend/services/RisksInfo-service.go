package services

import (
	"time"
)

// RisksInfoRequest servis modeli
type RisksInfoRequest struct {
	CampaignId  uint      `json:"campaign_id" validate:"required"` // Foreign key
	ProjectRisk string    `json:"project_risk" validate:"required"`
	SectorRisk  string    `json:"sector_risk" validate:"required"`
	ShareRisk   string    `json:"share_risk" validate:"required"`
	OtherRisk   string    `json:"other_risk" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Bu yapı, RisksInfo ile ilgili verileri işlemek için kullanılır.
