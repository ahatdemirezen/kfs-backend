package services

type VentureCategoryRequest struct {
	CampaignId uint   `json:"campaign_id" validate:"required"`
	Category   string `json:"category" validate:"required"`
}