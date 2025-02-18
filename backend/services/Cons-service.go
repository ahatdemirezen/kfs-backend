package services

import (
	"time"
)

// ConsRequest servis modeli
type ConsRequest struct {
	ConId          uint      `json:"con_id"`
	AnalysisInfoId uint      `json:"analysis_info_id" validate:"required"` // Foreign key
	Con            string    `json:"con" validate:"required"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// Bu yapı, Cons ile ilgili verileri işlemek için kullanılır.
