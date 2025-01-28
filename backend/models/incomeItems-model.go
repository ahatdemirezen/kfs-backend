package models

import (
	"time"
)

// IncomeItems modeli
type IncomeItems struct {
	IncomeItemId uint      `gorm:"primaryKey;autoIncrement;column:id"`  // Primary key
	CampaignId   uint      `gorm:"not null;column:campaign_id"`         // Foreign key
	Title        string    `gorm:"type:text;not null;column:title"`     // Gelir başlığı
	SalePrice    int       `gorm:"type:int;not null;column:sale_price"` // Satış fiyatı
	Cost         int       `gorm:"type:int;not null;column:cost"`       // Maliyet
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (IncomeItems) TableName() string {
	return "income_items"
}
