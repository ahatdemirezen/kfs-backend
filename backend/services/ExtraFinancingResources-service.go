package services

import (
	"time"
)

// ExtraFinancingResourceRequest servis modeli
type ExtraFinancingResourceRequest struct {
	FundingInfoId     uint      `json:"funding_info_id" validate:"required"` // Foreign key
	Description       string    `json:"description" validate:"required"`      // Ek finans kaynağı açıklaması
	SupplyDate        time.Time `json:"supply_date" validate:"required"`      // Tedarik tarihi
	Amount            int       `json:"amount" validate:"required"`           // Ek kaynak miktarı
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}


