package models

import (
	"time"
)

// OtherMarketTopic modeli
type OtherMarketTopic struct {
	TopicId      uint      `gorm:"primaryKey;autoIncrement;column:topic_id"`
	MarketInfoId uint      `gorm:"not null;column:market_info_id"` // Foreign key
	DocumentKey  string    `gorm:"type:text;column:document_key"`
	Subject      string    `gorm:"type:text;column:subject"`
	Description  string    `gorm:"type:text;column:description"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// MarketInfo tablosu ile ilişkilendirme
	MarketInfo MarketInfo `gorm:"foreignKey:MarketInfoId;references:MarketInfoId"`
}

// TableName tablosunun adını belirtir
func (OtherMarketTopic) TableName() string {
	return "other_market_topics"
}
