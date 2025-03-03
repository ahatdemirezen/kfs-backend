package models

import (
	"time"
)

// FinancialCategory modeli
type FinancialCategory struct {
	CategoryId   uint      `gorm:"primaryKey;autoIncrement;column:category_id"` // Primary key
	Category     string    `gorm:"type:text;column:category"`                   // Kategori adı
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"` // Oluşturulma zamanı
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"` // Güncellenme zamanı
}

// TableName tablosunun adını belirtir
func (FinancialCategory) TableName() string {
	return "financial_categories"
}
