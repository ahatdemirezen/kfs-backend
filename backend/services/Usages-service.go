package services

import (
	"time"
)

// UsageRequest servis modeli
type UsageRequest struct {
	FundingInfoId uint      `json:"funding_info_id" validate:"required"` // Foreign key
	Description   string    `json:"description" validate:"required"`      // Kullanım açıklaması
	StartingDate  time.Time `json:"starting_date" validate:"required"`   // Başlangıç tarihi
	EndingDate    time.Time `json:"ending_date" validate:"required"`     // Bitiş tarihi
	Amount        int       `json:"amount" validate:"required"`          // Kullanılan miktar
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// Bu yapı, Usage ile ilgili verileri işlemek için kullanılır.
