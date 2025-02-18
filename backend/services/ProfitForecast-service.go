package services

import (
	"time"
)

// ProfitForecastRequest service model
type ProfitForecastRequest struct {
	CampaignId     uint      `json:"campaign_id" validate:"required"` // Foreign key
	ProfitForecast string    `json:"profit_forecast" validate:"required"` // Profit forecast
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

