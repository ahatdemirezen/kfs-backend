package models

import (
	"time"
)

// Usages modeli
type Usage struct {
	UsageId       uint      `gorm:"primaryKey;autoIncrement;column:usage_id"`
	FundingInfoId uint      `gorm:"not null;column:funding_info_id"`     // Foreign key
	Description   string    `gorm:"type:text;column:description"`        // Kullanım açıklaması
	StartingDate  time.Time `gorm:"type:timestamp;column:starting_date"` // Başlangıç tarihi
	EndingDate    time.Time `gorm:"type:timestamp;column:ending_date"`   // Bitiş tarihi
	Amount        int       `gorm:"not null;column:amount"`              // Kullanılan miktar
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// İlişkiler
	FundingInfo FundingInfo `gorm:"foreignKey:FundingInfoId;references:FundingInfoId"`
}

// TableName tablosunun adını belirtir
func (Usage) TableName() string {
	return "usages"
}
