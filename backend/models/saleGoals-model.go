package models

import (
	"time"
)

// SaleGoals modeli
type SaleGoals struct {
	SaleGoalId   uint      `gorm:"primaryKey;autoIncrement;column:id"`  // Primary key
	CampaignId   uint      `gorm:"not null;column:campaign_id"`         // Campaign tablosundan foreign key
	IncomeItemId uint      `gorm:"not null;column:income_item_id"`      // IncomeItems tablosundan foreign key
	YearOne      int       `gorm:"type:int;not null;column:year_one"`   // 1. yıl hedefi
	YearTwo      int       `gorm:"type:int;not null;column:year_two"`   // 2. yıl hedefi
	YearThree    int       `gorm:"type:int;not null;column:year_three"` // 3. yıl hedefi
	YearFour     int       `gorm:"type:int;not null;column:year_four"`  // 4. yıl hedefi
	YearFive     int       `gorm:"type:int;not null;column:year_five"`  // 5. yıl hedefi
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`

	// IncomeItems tablosu ile ilişkilendirme
	IncomeItem IncomeItems `gorm:"foreignKey:IncomeItemId;references:IncomeItemId"`
}

// TableName tablosunun adını belirtir
func (SaleGoals) TableName() string {
	return "sale_goals"
}
