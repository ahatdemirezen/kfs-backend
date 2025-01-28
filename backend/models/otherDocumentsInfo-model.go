package models

import (
	"time"
)

// OtherDocumentsInfo modeli
type OtherDocumentsInfo struct {
	DocumentId  uint      `gorm:"primaryKey;autoIncrement;column:id"`     // Primary key
	CampaignId  uint      `gorm:"not null;column:campaign_id"`            // Foreign key
	DocumentKey string    `gorm:"type:text;not null;column:document_key"` // Belge anahtarı
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (OtherDocumentsInfo) TableName() string {
	return "other_documents_info"
}
