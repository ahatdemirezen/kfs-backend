package services

import (
	"time"
)

// VideosRequest service model
type VideosRequest struct {  
	VisualInfoId uint      `json:"visual_info_id" validate:"required"` // Foreign key
	VideoUrl     string    `json:"video_url" validate:"required"`      // Video URL
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}


