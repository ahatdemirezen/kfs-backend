package models

import (
	"time"
)

// Explanations modeli
type Explanations struct {
	ExplanationId uint      `gorm:"primaryKey;autoIncrement;column:id"`    // Primary key
	CampaignId    uint      `gorm:"not null;column:campaign_id"`           // Foreign key
	Explanation   string    `gorm:"type:text;not null;column:explanation"` // Açıklama metni
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (Explanations) TableName() string {
	return "explanations"
}
