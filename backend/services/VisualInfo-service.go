package services

import (
	"time"
)

// VisualInfoRequest service model
type VisualInfoRequest struct {
	CampaignId       uint      `json:"campaign_id" validate:"required"` // Foreign key
	ShowcasePhotoKey string    `json:"showcase_photo_key" validate:"required"` // Photo key
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}



