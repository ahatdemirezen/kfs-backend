package services

type PermissionRequest struct {
    CampaignId   uint   `json:"campaign_id" validate:"required"`
    DocumentKey  string `json:"document_key" validate:"required"`
    Subject      string `json:"subject" validate:"required"`
    Description  string `json:"description" validate:"required"`
}
