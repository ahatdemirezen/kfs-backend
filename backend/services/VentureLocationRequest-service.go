package services

type VentureLocationRequest struct {
	CampaignId uint   `json:"campaign_id" validate:"required"`
	Location   string `json:"location" validate:"required"`
}