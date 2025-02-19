package services

import (
	"time"
)

// ProsRequest servis modeli
type ProsRequest struct {
	AnalysisInfoId uint      `json:"analysis_info_id" validate:"required"` // Foreign key
	Pro            string    `json:"pro" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Bu yapı, Pros ile ilgili verileri işlemek için kullanılır.
