package services

import (
	"time"
)

// OtherProductTopicRequest servis modeli
type OtherProductTopicRequest struct {
	TopicId             uint      `json:"topic_id"`
	ProductModelInfoId  uint      `json:"product_model_info_id" validate:"required"` // Foreign key
	DocumentKey         string    `json:"document_key" validate:"required"`
	Subject             string    `json:"subject" validate:"required"`
	Description         string    `json:"description" validate:"required"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
