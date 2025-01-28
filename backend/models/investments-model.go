package models

import (
	"time"
)

// Investment model
type Investment struct {
	InvestmentId uint      `gorm:"primaryKey;autoIncrement;column:investment_id"`
	UserId       uint      `gorm:"not null;column:user_id"`              // Foreign key
	CampaignId   uint      `gorm:"not null;column:campaign_id"`          // Foreign key
	Balance      float64   `gorm:"type:numeric;not null;column:balance"` // Yatırım miktarı - Ondalık olabileceği için float64 kullanıldı (100,25 vs.)
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`

	// İlişkiler
	User     User     `gorm:"foreignKey:UserId;references:UserId"`
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (Investment) TableName() string {
	return "investments"
}
