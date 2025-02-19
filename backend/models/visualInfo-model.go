package models

import (
	"time"
)

// VisualInfo modeli
type VisualInfo struct {
	VisualInfoId     uint      `gorm:"primaryKey;autoIncrement;column:visual_info_id"` // Primary key
	CampaignId 		 uint      	`gorm:"not null;column:campaign_id"` // Foreign key
	ShowcasePhotoKey string    `gorm:"type:text;not null;column:showcase_photo_key"`   // Fotoğraf anahtarı
	CreatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt        time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (VisualInfo) TableName() string {
	return "visual_info"
}
