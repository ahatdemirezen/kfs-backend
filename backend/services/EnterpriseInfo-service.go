package services

import (
	"time"
)

// EnterpriseInfoRequest servis modeli
type EnterpriseInfoRequest struct {
	CampaignId        uint      `json:"campaign_id" validate:"required"` // Foreign key
	EnterpriseName    string    `json:"enterprise_name" validate:"required"`    // Girişim adı
	EnterpriseCapital int       `json:"enterprise_capital" validate:"required"` // Sermaye miktarı
	EnterpriseCity    string    `json:"enterprise_city" validate:"required"`    // Şehir
	EnterpriseTown    string    `json:"enterprise_town" validate:"required"`    // İlçe
	EnterpriseAddress string    `json:"enterprise_address" validate:"required"` // Adres
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

