package models

import (
	"time"
)

// Permission modeli
type Permission struct {
	PermissionId uint      `gorm:"primaryKey;autoIncrement;column:permission_id"`
	CampaignId   uint      `gorm:"not null;column:campaign_id"` // Foreign key
	DocumentKey  string    `gorm:"type:text;column:document_key"`
	Subject      string    `gorm:"type:text;column:subject"`
	Description  string    `gorm:"type:text;column:description"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (Permission) TableName() string {
	return "permissions"
}
