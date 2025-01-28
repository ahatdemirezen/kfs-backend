package models

import (
	"time"
)

// ExtraFinancingResources modeli
type ExtraFinancingResource struct {
	FinanceResourceId uint      `gorm:"primaryKey;autoIncrement;column:finance_resource_id"`
	FundingInfoId     uint      `gorm:"not null;column:funding_info_id"`   // Foreign key
	Description       string    `gorm:"type:text;column:description"`      // Ek finans kaynağı açıklaması
	SupplyDate        time.Time `gorm:"type:timestamp;column:supply_date"` // Tedarik tarihi
	Amount            int       `gorm:"not null;column:amount"`            // Ek kaynak miktarı
	CreatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// İlişkiler
	FundingInfo FundingInfo `gorm:"foreignKey:FundingInfoId;references:FundingInfoId"`
}

// TableName tablosunun adını belirtir
func (ExtraFinancingResource) TableName() string {
	return "extra_financing_resources"
}
