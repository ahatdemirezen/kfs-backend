package models

import (
	"time"
)

// ParticipantEmail modeli
type ParticipantEmail struct {
	ParticipantId uint      `gorm:"primaryKey;autoIncrement;column:participant_id"`
	CampaignId    uint      `gorm:"not null;column:campaign_id"` // Foreign key
	Email         string    `gorm:"type:text;not null;unique;column:email"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (ParticipantEmail) TableName() string {
	return "participant_emails"
}
