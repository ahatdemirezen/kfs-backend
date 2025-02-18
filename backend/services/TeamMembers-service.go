package services

import (
    "time"
)

// TeamMemberRequest servis modeli
type TeamMemberRequest struct {
    CampaignId     uint      `json:"campaign_id" validate:"required"` // Foreign key
    Name           string    `json:"name" validate:"required"`
    Surname        string    `json:"surname" validate:"required"`
    Position       string    `json:"position" validate:"required"`
    ResumeKey      string    `json:"resume_key" validate:"required"`
    PhotoKey       string    `json:"photo_key" validate:"required"`
    Biography      string    `json:"biography" validate:"required"`
    Responsibility string    `json:"responsibility" validate:"required"`
    Profession     string    `json:"profession" validate:"required"`
    Relation       string    `json:"relation" validate:"required"`
    Email          string    `json:"email" validate:"required"`
    Instagram      string    `json:"instagram" validate:"required"`
    Twitter        string    `json:"twitter" validate:"required"`
    Linkedin       string    `json:"linkedin" validate:"required"`
    CreatedAt      time.Time `json:"created_at" validate:"required"`
    UpdatedAt      time.Time `json:"updated_at" validate:"required"`
}

// Bu yapı, TeamMember ile ilgili verileri işlemek için kullanılır.
