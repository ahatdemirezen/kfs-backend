package models

import (
	"time"
)

// SocialMedia modeli
type SocialMedia struct {
	SocialId      uint      `gorm:"primaryKey;autoIncrement;column:id"` // Primary key
	ProfileId     uint      `gorm:"not null;column:profile_id"`         // Foreign key
	URL           string    `gorm:"type:text;not null;column:url"`      // Sosyal medya URL'si
	FollowerCount string    `gorm:"type:text;column:follower_count"`    // Takipçi sayısı
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Profile tablosu ile ilişkilendirme
	Profile Profile `gorm:"foreignKey:ProfileId;references:ProfileId"`
}

// TableName tablosunun adını belirtir
func (SocialMedia) TableName() string {
	return "social_media"
}
