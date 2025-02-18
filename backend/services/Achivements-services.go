package services

import (
    "time"  // Import the time package to use time.Time
)

type AchievementRequest struct {
    CampaignId   uint      `json:"campaign_id" validate:"required"`
    Date         time.Time `json:"date" validate:"required"`
    Foundation   string    `json:"foundation" validate:"required"`
    Description  string    `json:"description" validate:"required"`
    DocumentKey  string    `json:"document_key" validate:"required"`
}
