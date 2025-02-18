package services

type ParticipantEmailRequest struct {
    CampaignId uint   `json:"campaign_id" validate:"required"`
    Email      string `json:"email" validate:"required,email"`
}
