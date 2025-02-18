package services

import (
	"time"
)

// SaleGoalsRequest service model
type SaleGoalsRequest struct {
	CampaignId   uint      `json:"campaign_id" validate:"required"` // Campaign foreign key
	IncomeItemId uint      `json:"income_item_id" validate:"required"` // IncomeItems foreign key
	YearOne      int       `json:"year_one" validate:"required"`
	YearTwo      int       `json:"year_two" validate:"required"`
	YearThree    int       `json:"year_three" validate:"required"`
	YearFour     int       `json:"year_four" validate:"required"`
	YearFive     int       `json:"year_five" validate:"required"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
