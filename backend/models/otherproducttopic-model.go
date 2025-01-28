package models

import (
	"time"
)

// OtherProductTopic modeli
type OtherProductTopic struct {
	TopicId           uint      `gorm:"primaryKey;autoIncrement;column:topic_id"`
	ProductModelInfoId uint      `gorm:"not null;column:product_model_info_id"` // Foreign key
	DocumentKey       string    `gorm:"type:text;column:document_key"`
	Subject           string    `gorm:"type:text;column:subject"`
	Description       string    `gorm:"type:text;column:description"`
	CreatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// ProductModelInfo tablosu ile ilişkilendirme
	ProductModelInfo ProductModelInfo `gorm:"foreignKey:ProductModelInfoId;references:ProductModelInfoId"`
}

// TableName tablosunun adını belirtir
func (OtherProductTopic) TableName() string {
	return "other_product_topics"
}
