package services

import (
	"time"
)

// FundingInfoRequest servis modeli
type FundingInfoRequest struct {
	CampaignId           uint      `json:"campaign_id" validate:"required"` // Foreign key
	VentureValue         int       `json:"venture_value" validate:"required"` // Girişimin toplam değeri
	RequiredVentureFund  int       `json:"required_venture_fund" validate:"required"` // İhtiyaç duyulan fon
	FundingMonths        int       `json:"funding_months" validate:"required"` // Fonun geçerli olduğu ay sayısı
	EvaluationReportKey  string    `json:"evaluation_report_key" validate:"required"` // Değerlendirme raporu
	SharePercentage      int       `json:"share_percentage" validate:"required"` // Pay yüzdesi
	ExtraFunding         bool      `json:"extra_funding"` // Ek fon gereksinimi
	ComparingPartnership string    `json:"comparing_partnership"` // Karşılaştırmalı ortaklık bilgisi
	GeneralReason        string    `json:"general_reason"` // Genel açıklama/yorum
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

