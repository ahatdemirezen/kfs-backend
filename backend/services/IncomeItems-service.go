package services

import (
	"time"
)

// IncomeItemsRequest service model
type IncomeItemsRequest struct {
	CampaignId   uint      `json:"campaign_id" validate:"required"` // Foreign key
	Title        string    `json:"title" validate:"required"`       // Revenue item title
	SalePrice    int       `json:"sale_price" validate:"required"`  // Sale price
	Cost         int       `json:"cost" validate:"required"`        // Cost
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
