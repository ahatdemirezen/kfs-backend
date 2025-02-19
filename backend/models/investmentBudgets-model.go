package models

import (
	"time"
)

// InvestmentBudgets modeli
type InvestmentBudgets struct {
	BudgetsId       uint      `gorm:"primaryKey;autoIncrement;column:budgets_id"` // Primary key
	CampaignId      uint      `gorm:"not null;column:campaign_id"`                // Foreign key
	YearOneBudget   int       `gorm:"type:int;not null;column:year_one_budget"`   // 1. yıl bütçesi
	YearTwoBudget   int       `gorm:"type:int;not null;column:year_two_budget"`   // 2. yıl bütçesi
	YearThreeBudget int       `gorm:"type:int;not null;column:year_three_budget"` // 3. yıl bütçesi
	YearFourBudget  int       `gorm:"type:int;not null;column:year_four_budget"`  // 4. yıl bütçesi
	YearFiveBudget  int       `gorm:"type:int;not null;column:year_five_budget"`  // 5. yıl bütçesi
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (InvestmentBudgets) TableName() string {
	return "investment_budgets"
}
