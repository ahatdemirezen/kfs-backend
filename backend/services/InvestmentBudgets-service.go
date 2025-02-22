package services

import (
	"time"
)

// InvestmentBudgetsRequest service model
type InvestmentBudgetsRequest struct {
	CampaignId      uint      `json:"campaign_id" validate:"required"` // Foreign key
	YearOneBudget   int       `json:"year_one_budget" validate:"required"`
	YearTwoBudget   int       `json:"year_two_budget" validate:"required"`
	YearThreeBudget int       `json:"year_three_budget" validate:"required"`
	YearFourBudget  int       `json:"year_four_budget" validate:"required"`
	YearFiveBudget  int       `json:"year_five_budget" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
