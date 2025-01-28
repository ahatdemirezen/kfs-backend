package models

import (
	"time"
)

// MarketInfo modeli
type MarketInfo struct {
	MarketInfoId              uint      `gorm:"primaryKey;autoIncrement;column:market_info_id"`
	CampaignId                uint      `gorm:"not null;column:campaign_id"` // Foreign key
	AboutMarket               string    `gorm:"type:text;column:about_market"`
	AboutCompetition          string    `gorm:"type:text;column:about_competition"`
	TargetSummary             string    `gorm:"type:text;column:target_summary"`
	CommercializationSummary  string    `gorm:"type:text;column:commercialization_summary"`
	AboutMarketKey            string    `gorm:"type:text;column:about_market_key"`
	AboutCompetitionKey       string    `gorm:"type:text;column:about_competition_key"`
	TargetSummaryKey          string    `gorm:"type:text;column:target_summary_key"`
	CommercializationSummaryKey string  `gorm:"type:text;column:commercialization_summary_key"`
	CreatedAt                 time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt                 time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (MarketInfo) TableName() string {
	return "market_info"
}
