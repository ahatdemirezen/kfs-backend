package models

import (
	"time"
)

// RisksInfo modeli
type RisksInfo struct {
	RisksInfoId uint      `gorm:"primaryKey;autoIncrement;column:risks_info_id"`
	CampaignId  uint      `gorm:"not null;column:campaign_id"` // Foreign key
	ProjectRisk string    `gorm:"type:text;column:project_risk"`
	SectorRisk  string    `gorm:"type:text;column:sector_risk"`
	ShareRisk   string    `gorm:"type:text;column:share_risk"`
	OtherRisk   string    `gorm:"type:text;column:other_risk"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (RisksInfo) TableName() string {
	return "risks_info"
}
