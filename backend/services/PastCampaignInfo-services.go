package services

type PastCampaignInfoRequest struct {
    CampaignId  uint   `json:"campaign_id" validate:"required"`
    Status      bool   `json:"status" validate:"required"`
    Description string `json:"description" validate:"required"`
}
