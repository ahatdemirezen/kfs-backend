package models

import (
	"time"
)

// AfterFundingFounderPartners modeli
type AfterFundingFounderPartner struct {
	PartnerId              uint      `gorm:"primaryKey;autoIncrement;column:partner_id"`
	EnterpriseInfoId       uint      `gorm:"not null;column:enterprise_info_id"`        // Foreign key
	PartnerName            string    `gorm:"type:text;not null;column:partner_name"`    // Ortak adı
	PartnerSurname         string    `gorm:"type:text;not null;column:partner_surname"` // Ortak soyadı
	PartnerTitle           string    `gorm:"type:text;column:partner_title"`            // Ortak unvanı
	PartnerSchool          string    `gorm:"type:text;column:partner_school"`           // Mezun olduğu okul
	PartnerGpa             int       `gorm:"column:partner_gpa"`                        // Mezuniyet not ortalaması
	ResumeKey              string    `gorm:"type:text;column:resume_key"`               // Özgeçmiş dosya anahtarı
	Citizenship            string    `gorm:"type:text;column:citizenship"`              // Vatandaşlık bilgisi
	CapitalShareAmount     int       `gorm:"not null;column:capital_share_amount"`      // Sermaye payı miktarı
	CapitalSharePercentage int       `gorm:"not null;column:capital_share_percentage"`  // Sermaye payı yüzdesi
	VotePercentage         int       `gorm:"not null;column:vote_percentage"`           // Oy hakkı yüzdesi
	Privilege              string    `gorm:"type:text;column:privilege"`                // Ayrıcalıklar
	CampaignRelation       string    `gorm:"type:text;column:campaign_relation"`        // Kampanyayla ilişkisi
	Experience             string    `gorm:"type:text;column:experience"`               // Deneyim bilgisi
	Profession             string    `gorm:"type:text;column:profession"`               // Mesleği
	CreatedAt              time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt              time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// İlişkiler
	EnterpriseInfo EnterpriseInfo `gorm:"foreignKey:EnterpriseInfoId;references:EnterpriseInfoId"`
}

// TableName tablosunun adını belirtir
func (AfterFundingFounderPartner) TableName() string {
	return "after_funding_founder_partners"
}
