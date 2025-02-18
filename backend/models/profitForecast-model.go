package models

import (
	"time"
)

// ProfitForecast modeli
type ProfitForecast struct {
	ForecastId     uint      `gorm:"primaryKey;autoIncrement;column:forecast_id"` // Primary key
	CampaignId     uint      `gorm:"not null;column:campaign_id"`                 // Foreign key
	ProfitForecast string    `gorm:"type:text;not null;column:profit_forecast"`   // Profit forecast
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (ProfitForecast) TableName() string {
	return "profit_forecast"
}
