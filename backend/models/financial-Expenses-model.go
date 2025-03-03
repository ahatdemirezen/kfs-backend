package models

import (
	"time"
)

// FinancialExpense modeli
type FinancialExpense struct {
	ExpenseId     uint      `gorm:"primaryKey;autoIncrement;column:expense_id"` // Primary key
	CampaignId    uint      `gorm:"not null;index;column:campaign_id"`          // Foreign key (Kampanya)
	Year          int       `gorm:"column:year"`                                // Harcamanın yapıldığı yıl
	SubCategoryId uint      `gorm:"not null;index;column:sub_category_id"`      // Foreign key (Alt kategori)
	Value         int       `gorm:"column:value"`                               // Harcama değeri
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"` // Oluşturulma zamanı
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"` // Güncellenme zamanı

	// İlişkilendirme
	Campaign        Campaign            `gorm:"foreignKey:CampaignId;references:CampaignId"`
	FinancialSubCategory FinancialSubCategory `gorm:"foreignKey:SubCategoryId;references:SubCategoryId"`
}

// TableName tablosunun adını belirtir
func (FinancialExpense) TableName() string {
	return "financial_expenses"
}
