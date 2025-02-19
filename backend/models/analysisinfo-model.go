package models

import (
	"time"
)

// AnalysisInfo modeli
type AnalysisInfo struct {
	AnalysisInfoId uint      `gorm:"primaryKey;autoIncrement;column:analysis_info_id"`
	CampaignId     uint      `gorm:"not null;column:campaign_id"` // Foreign key
	SwotKey        string    `gorm:"type:text;column:swot_key"`
	BusinessKey    string    `gorm:"type:text;column:business_key"`
	InvestorKey    string    `gorm:"type:text;column:investor_key"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (AnalysisInfo) TableName() string {
	return "analysis_info"
}
