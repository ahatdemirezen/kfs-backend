package services

type VentureBusinessModalRequest struct {
    CampaignId      uint   `json:"campaign_id" validate:"required"`
    BusinessModal   string `json:"business_modal" validate:"required"`
}
