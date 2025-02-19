package models

import (
	"time"
)

// FundingInfo modeli
type FundingInfo struct {
	FundingInfoId        uint      `gorm:"primaryKey;autoIncrement;column:funding_info_id"`
	CampaignId           uint      `gorm:"not null;column:campaign_id"`            // Foreign key
	VentureValue         int       `gorm:"not null;column:venture_value"`          // Girişimin toplam değeri
	RequiredVentureFund  int       `gorm:"not null;column:required_venture_fund"`  // İhtiyaç duyulan fon
	FundingMonths        int       `gorm:"not null;column:funding_months"`         // Fonun geçerli olduğu ay sayısı
	EvaluationReportKey  string    `gorm:"type:text;column:evaluation_report_key"` // Değerlendirme raporu
	SharePercentage      int       `gorm:"not null;column:share_percentage"`       // Pay yüzdesi
	ExtraFunding         *bool      `gorm:"type:boolean;not null;column:extra_funding"`     // Ek fon gereksinimi
	ComparingPartnership string    `gorm:"type:text;column:comparing_partnership"` // Karşılaştırmalı ortaklık bilgisi
	GeneralReason        string    `gorm:"type:text;column:general_reason"`        // Genel açıklama/yorum
	CreatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt            time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// İlişkiler
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId" json:"-"`
}

// TableName tablosunun adını belirtir
func (FundingInfo) TableName() string {
	return "funding_info"
}
