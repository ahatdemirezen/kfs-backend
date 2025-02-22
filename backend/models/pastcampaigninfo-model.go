package models

import (
	"time"
)

// PastCampaignInfo modeli
type PastCampaignInfo struct {
	PastCampaignInfoId uint      `gorm:"primaryKey;autoIncrement;column:past_campaign_info_id"`
	CampaignId         uint      `gorm:"not null;column:campaign_id"` 
	Status             *bool     `gorm:"type:boolean;not null;column:status"` 
	Description        string    `gorm:"type:text;column:description"`
	CreatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt          time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (PastCampaignInfo) TableName() string {
	return "past_campaign_infos"
}
