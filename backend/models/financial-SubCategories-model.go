package models

import (
	"time"
)

// FinancialSubCategory modeli
type FinancialSubCategory struct {
	SubCategoryId uint      `gorm:"primaryKey;autoIncrement;column:sub_category_id"` // Primary key
	CategoryId    uint      `gorm:"not null;index;column:category_id"`               // Foreign key
	SubCategory   string    `gorm:"type:text;column:sub_category"`                   // Alt kategori adı
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"` // Oluşturulma zamanı
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"` // Güncellenme zamanı

	// Category ile ilişki
	Category FinancialCategory `gorm:"foreignKey:CategoryId;references:CategoryId"`
}

// TableName tablosunun adını belirtir
func (FinancialSubCategory) TableName() string {
	return "financial_sub_categories"
}
