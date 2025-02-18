package services

type VentureSectorRequest struct {
    CampaignId uint   `json:"campaign_id" validate:"required"`
    Sector     string `json:"sector" validate:"required"`
}
