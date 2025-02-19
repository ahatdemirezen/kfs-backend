package models

import (
	"time"
)

// Cons modeli
type Cons struct {
	ConId          uint      `gorm:"primaryKey;autoIncrement;column:con_id"`
	AnalysisInfoId uint      `gorm:"not null;column:analysis_info_id"` // Foreign key
	Con            string    `gorm:"type:text;column:con"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// AnalysisInfo tablosu ile ilişkilendirme
	AnalysisInfo AnalysisInfo `gorm:"foreignKey:AnalysisInfoId;references:AnalysisInfoId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Cons) TableName() string {
	return "cons"
}
