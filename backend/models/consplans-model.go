package models

import (
	"time"
)

// ConsPlan modeli
type ConsPlan struct {
	ConPlanId      uint      `gorm:"primaryKey;autoIncrement;column:con_plan_id"`
	AnalysisInfoId uint      `gorm:"not null;column:analysis_info_id"` // Foreign key
	ConPlan        string    `gorm:"type:text;column:con_plan"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// AnalysisInfo tablosu ile ilişkilendirme
	AnalysisInfo AnalysisInfo `gorm:"foreignKey:AnalysisInfoId;references:AnalysisInfoId"`
}

// TableName tablosunun adını belirtir
func (ConsPlan) TableName() string {
	return "cons_plans"
}
