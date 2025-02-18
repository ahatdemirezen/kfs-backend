package models

import (
	"time"
)

// VentureLocation modeli
type VentureLocation struct {
	LocationId uint      `gorm:"primaryKey;autoIncrement;column:location_id"`
	CampaignId uint      `gorm:"not null;column:campaign_id"` // Foreign key
	Location   string    `gorm:"type:text;column:location"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (VentureLocation) TableName() string {
	return "venture_locations"
}
