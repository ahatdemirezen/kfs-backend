package models

import (
	"time"
)

// Videos modeli
type Videos struct {
	VideoId      uint      `gorm:"primaryKey;autoIncrement;column:video_id"`  // Primary key
	VisualInfoId uint      `gorm:"not null;column:visual_info_id"`      // Foreign key
	VideoUrl     string    `gorm:"type:text;not null;column:video_url"` // Video URL
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// VisualInfo tablosu ile ilişkilendirme
	VisualInfo VisualInfo `gorm:"foreignKey:VisualInfoId;references:VisualInfoId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Videos) TableName() string {
	return "videos"
}
