package models

import (
	"time"
)

// Patent modeli
type Patent struct {
	PatentId       uint      `gorm:"primaryKey;autoIncrement;column:patent_id"`
	CampaignId     uint      `gorm:"not null;column:campaign_id"` // Foreign key
	DocumentKey    string    `gorm:"type:text;column:document_key"`
	DocumentNumber string    `gorm:"type:text;column:document_number"`
	Description    string    `gorm:"type:text;column:description"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Patent) TableName() string {
	return "patents"
}
