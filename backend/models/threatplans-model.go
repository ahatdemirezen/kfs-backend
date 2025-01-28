package models

import (
	"time"
)

// ThreatPlan modeli
type ThreatPlan struct {
	ThreatPlanId   uint      `gorm:"primaryKey;autoIncrement;column:threat_plan_id"`
	AnalysisInfoId uint      `gorm:"not null;column:analysis_info_id"` // Foreign key
	ThreatPlan     string    `gorm:"type:text;column:threat_plan"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// AnalysisInfo tablosu ile ilişkilendirme
	AnalysisInfo AnalysisInfo `gorm:"foreignKey:AnalysisInfoId;references:AnalysisInfoId"`
}

// TableName tablosunun adını belirtir
func (ThreatPlan) TableName() string {
	return "threat_plans"
}
