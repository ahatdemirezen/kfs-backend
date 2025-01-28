package models

import (
	"time"
)

// FinancialDocuments modeli
type FinancialDocuments struct {
	FinancialDocumentId uint      `gorm:"primaryKey;autoIncrement;column:id"`             // Primary key
	CampaignId          uint      `gorm:"not null;column:campaign_id"`                    // Foreign key
	Subject             string    `gorm:"type:text;not null;column:subject"`              // Konu başlığı
	DocumentKey         string    `gorm:"type:varchar(255);not null;column:document_key"` // Belge anahtarı
	CreatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (FinancialDocuments) TableName() string {
	return "financial_documents"
}
