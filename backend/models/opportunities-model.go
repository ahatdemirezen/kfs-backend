package models

import (
	"time"
)

// Opportunity modeli
type Opportunity struct {
	OpportunityId  uint      `gorm:"primaryKey;autoIncrement;column:opportunity_id"`
	AnalysisInfoId uint      `gorm:"not null;column:analysis_info_id"` // Foreign key
	Opportunity    string    `gorm:"type:text;column:opportunity"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// AnalysisInfo tablosu ile ilişkilendirme
	AnalysisInfo AnalysisInfo `gorm:"foreignKey:AnalysisInfoId;references:AnalysisInfoId"`
}

// TableName tablosunun adını belirtir
func (Opportunity) TableName() string {
	return "opportunities"
}
