package models

import (
	"time"
)

// VentureBusinessModal modeli
type VentureBusinessModal struct {
	BusinessModalId uint      `gorm:"primaryKey;autoIncrement;column:business_modal_id"`
	CampaignId      uint      `gorm:"not null;column:campaign_id"` // Foreign key
	BusinessModal   string    `gorm:"type:text;column:business_modal"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (VentureBusinessModal) TableName() string {
	return "venture_business_modals"
}
