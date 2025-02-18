package models

import (
	"time"
)

// VentureSector modeli
type VentureSector struct {
	SectorId   uint      `gorm:"primaryKey;autoIncrement;column:sector_id"`
	CampaignId uint      `gorm:"not null;column:campaign_id"` // Foreign key
	Sector     string    `gorm:"type:text;column:sector"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (VentureSector) TableName() string {
	return "venture_sectors"
}
