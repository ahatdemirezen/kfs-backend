package services

type PatentRequest struct {
    CampaignId     uint   `json:"campaign_id" validate:"required"`
    DocumentKey    string `json:"document_key" validate:"required"`
    DocumentNumber string `json:"document_number" validate:"required"`
    Description    string `json:"description" validate:"required"`
}
