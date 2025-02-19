package services

import (
	"time"
)

// OtherMarketTopicRequest servis modeli
type OtherMarketTopicRequest struct {
	MarketInfoId  uint      `json:"market_info_id" validate:"required"` // Foreign key
	DocumentKey   string    `json:"document_key" validate:"required"`
	Subject       string    `json:"subject" validate:"required"`
	Description   string    `json:"description" validate:"required"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

