package models

import (
	"time"
)

// Threat modeli
type Threat struct {
	ThreatId       uint      `gorm:"primaryKey;autoIncrement;column:threat_id"`
	AnalysisInfoId uint      `gorm:"not null;column:analysis_info_id"` // Foreign key
	Threat         string    `gorm:"type:text;column:threat"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// AnalysisInfo tablosu ile ilişkilendirme
	AnalysisInfo AnalysisInfo `gorm:"foreignKey:AnalysisInfoId;references:AnalysisInfoId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Threat) TableName() string {
	return "threats"
}
