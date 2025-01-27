package models

import (
	"time"
)

// Campaign modeli
type Campaign struct {
	CampaignId      uint      `gorm:"primaryKey;autoIncrement;column:campaign_id"`
	UserId          uint      `gorm:"not null;column:user_id"` // Foreign key
	CampaignStatus  string    `gorm:"type:text;column:campaign_status"`
	CampaignCode    string    `gorm:"type:text;column:campaign_code"`
	CampaignLogoKey string    `gorm:"type:text;column:campaign_logo_key"`
	VentureName     string    `gorm:"type:text;column:venture_name"`
	CampaignName    string    `gorm:"type:text;column:campaign_name"`
	Description     string    `gorm:"type:text;column:description"`
	AboutProject    string    `gorm:"type:text;column:about_project"`
	Summary         string    `gorm:"type:text;column:summary"`
	VenturePurpose  string    `gorm:"type:text;column:venture_purpose"`
	VenturePhase    string    `gorm:"type:text;column:venture_phase"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// User tablosu ile ilişkilendirme
	User User `gorm:"foreignKey:UserId;references:UserId"`
}

// TableName tablosunun adını belirtir
func (Campaign) TableName() string {
	return "campaigns"
}
