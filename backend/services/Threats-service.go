package services

import (
	"time"
)

// ThreatRequest servis modeli
type ThreatRequest struct {
	AnalysisInfoId uint      `json:"analysis_info_id" validate:"required"` // Foreign key
	Threat         string    `json:"threat" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Bu yapı, Threat ile ilgili verileri işlemek için kullanılır.
