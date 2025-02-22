package services

import (
	"time"
)

// ProductModelInfoRequest servis modeli
type ProductModelInfoRequest struct {
	CampaignId           uint      `json:"campaign_id" validate:"required"`
	ProductSummary       string    `json:"product_summary" validate:"required"`
	AboutProduct         string    `json:"about_product" validate:"required"`
	Problem              string    `json:"problem" validate:"required"`
	Solve                string    `json:"solve" validate:"required"`
	ValueProposition     string    `json:"value_proposition" validate:"required"`
	ProcessSummary       string    `json:"process_summary" validate:"required"`
	AboutProcess         string    `json:"about_process" validate:"required"`
	AboutSideProduct     string    `json:"about_side_product" validate:"required"`
	TechnicalAnalyses    string    `json:"technical_analyses" validate:"required"`
	ArgeSummary          string    `json:"arge_summary" validate:"required"`
	PreviousSales        string    `json:"previous_sales" validate:"required"`
	AboutProductKey      string    `json:"about_product_key" validate:"required"`
	ProcessSummaryKey    string    `json:"process_summary_key" validate:"required"`
	AboutProcessKey      string    `json:"about_process_key" validate:"required"`
	AboutSideKey         string    `json:"about_side_key" validate:"required"`
	TechnicalAnalysesKey string    `json:"technical_analyses_key" validate:"required"`
	ArgeSummaryKey       string    `json:"arge_summary_key" validate:"required"`
	PreviousSalesKey     string    `json:"previous_sales_key" validate:"required"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}