package models

import (
	"time"
)

// OtherPhotos modeli
type OtherPhotos struct {
	PhotoId      uint      `gorm:"primaryKey;autoIncrement;column:photo_id"` // Primary key
	VisualInfoId uint      `gorm:"not null;column:visual_info_id"`           // Foreign key
	PhotoKey     string    `gorm:"type:text;not null;column:photo_key"`      // Photo key
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// VisualInfo tablosu ile ilişkilendirme
	VisualInfo VisualInfo `gorm:"foreignKey:VisualInfoId;references:VisualInfoId" json:"-"`
}

// TableName tablosunun adını belirtir
func (OtherPhotos) TableName() string {
	return "other_photos"
}
