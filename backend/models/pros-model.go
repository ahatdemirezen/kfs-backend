package models

import (
	"time"
)

// Pros modeli
type Pros struct {
	ProId          uint      `gorm:"primaryKey;autoIncrement;column:pro_id"`
	AnalysisInfoId uint      `gorm:"not null;column:analysis_info_id"` // Foreign key
	Pro            string    `gorm:"type:text;column:pro"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// AnalysisInfo tablosu ile ilişkilendirme
	AnalysisInfo AnalysisInfo `gorm:"foreignKey:AnalysisInfoId;references:AnalysisInfoId"`
}

// TableName tablosunun adını belirtir
func (Pros) TableName() string {
	return "pros"
}
