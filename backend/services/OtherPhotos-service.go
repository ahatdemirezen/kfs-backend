package services

import (
	"time"
)

// OtherPhotosRequest service model
type OtherPhotosRequest struct {
	VisualInfoId uint      `json:"visual_info_id" validate:"required"` // Foreign key
	PhotoKey     string    `json:"photo_key" validate:"required"`      // Photo key
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}


