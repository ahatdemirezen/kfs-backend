package models

import (
	"time"
)

// Explanations modeli
type Explanations struct {
	ExplanationId uint      `gorm:"primaryKey;autoIncrement;column:explanation_id"` // Primary key
	CampaignId    uint      `gorm:"not null;column:campaign_id"`                    // Foreign key
	Explanation   string    `gorm:"type:text;not null;column:explanation"`          // Explanation text
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Explanations) TableName() string {
	return "explanations"
}
