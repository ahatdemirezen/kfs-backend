package services

import (
	"time"
)

// MarketInfoRequest servis modeli
type MarketInfoRequest struct {
	CampaignId                 uint      `json:"campaign_id" validate:"required"` // Foreign key
	AboutMarket                string    `json:"about_market" validate:"required"`
	AboutCompetition           string    `json:"about_competition" validate:"required"`
	TargetSummary              string    `json:"target_summary" validate:"required"`
	CommercializationSummary   string    `json:"commercialization_summary" validate:"required"`
	AboutMarketKey             string    `json:"about_market_key" validate:"required"`
	AboutCompetitionKey        string    `json:"about_competition_key" validate:"required"`
	TargetSummaryKey           string    `json:"target_summary_key" validate:"required"`
	CommercializationSummaryKey string   `json:"commercialization_summary_key" validate:"required"`
	CreatedAt                  time.Time `json:"created_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
}

// Bu yapı, MarketInfo ile ilgili verileri işlemek için kullanılır.
