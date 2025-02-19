package models

import (
	"time"
)

// Achievement modeli
type Achievement struct {
	AchievementId uint      `gorm:"primaryKey;autoIncrement;column:achievement_id"`
	CampaignId    uint      `gorm:"not null;column:campaign_id"` // Foreign key
	Date          time.Time `gorm:"type:timestamp;column:date"`
	Foundation    string    `gorm:"type:text;column:foundation"`
	Description   string    `gorm:"type:text;column:description"`
	DocumentKey   string    `gorm:"type:text;column:document_key"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Achievement) TableName() string {
	return "achievements"
}
