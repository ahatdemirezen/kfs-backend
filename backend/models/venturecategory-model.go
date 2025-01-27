package models

import (
	"time"
)

// VentureCategory modeli
type VentureCategory struct {
	CategoryId uint      `gorm:"primaryKey;autoIncrement;column:category_id"`
	CampaignId uint      `gorm:"not null;column:campaign_id"` // Foreign key
	Category   string    `gorm:"type:text;column:category"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (VentureCategory) TableName() string {
	return "venture_categories"
}
