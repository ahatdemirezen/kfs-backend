package models

import (
	"time"
)

// EnterpriseInfo modeli
type EnterpriseInfo struct {
	EnterpriseInfoId  uint      `gorm:"primaryKey;autoIncrement;column:enterprise_info_id"`
	CampaignId        uint      `gorm:"not null;column:campaign_id"`                  // Foreign key
	EnterpriseName    string    `gorm:"type:text;not null;column:enterprise_name"`    // Girişim adı
	EnterpriseCapital int       `gorm:"not null;column:enterprise_capital"`           // Sermaye miktarı
	EnterpriseCity    string    `gorm:"type:text;not null;column:enterprise_city"`    // Şehir
	EnterpriseTown    string    `gorm:"type:text;not null;column:enterprise_town"`    // İlçe
	EnterpriseAddress string    `gorm:"type:text;not null;column:enterprise_address"` // Adres
	CreatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt         time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// İlişkiler
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (EnterpriseInfo) TableName() string {
	return "enterprise_info"
}
