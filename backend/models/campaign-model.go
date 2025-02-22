package models

import (
	"time"
)

// Campaign modeli
type Campaign struct {
	CampaignId      uint      `gorm:"primaryKey;autoIncrement;column:campaign_id"`
	UserId          uint      `gorm:"not null;column:user_id"` // Foreign key
	CampaignStatus  string    `gorm:"type:text;column:campaign_status"`
	CampaignCode    string    `gorm:"type:text;column:campaign_code"`
	CampaignLogoKey string    `gorm:"type:text;column:campaign_logo_key"`
	VentureName     string    `gorm:"type:text;column:venture_name"`
	CampaignName    string    `gorm:"type:text;column:campaign_name"`
	Description     string    `gorm:"type:text;column:description"`
	AboutProject    string    `gorm:"type:text;column:about_project"`
	Summary         string    `gorm:"type:text;column:summary"`
	VenturePurpose  string    `gorm:"type:text;column:venture_purpose"`
	VenturePhase    string    `gorm:"type:text;column:venture_phase"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// User tablosu ile ilişkilendirme
	User User `gorm:"foreignKey:UserId;references:UserId" json:"-"`
}

// TableName tablosunun adını belirtir
func (Campaign) TableName() string {
	return "campaigns"
}


// package models

// import (
// 	"time"
// )

// // Campaign modeli
// type Campaign struct {
// 	ID              uint    `json:"id" gorm:"primaryKey"`
// 	UserId          uint    `json:"user_id,omitempty" gorm:"not null;column:user_id"` // Foreign key
// 	CampaignStatus  string  `json:"campaign_status,omitempty" gorm:"type:text;column:campaign_status"`
// 	CampaignCode    string  `json:"campaign_code,omitempty" gorm:"type:text;column:campaign_code"`
// 	CampaignLogoKey string  `json:"campaign_logo_key,omitempty" gorm:"type:text;column:campaign_logo_key"`
// 	VentureName     string  `json:"venture_name,omitempty" gorm:"type:text;column:venture_name"`
// 	CampaignName    string  `json:"campaign_name,omitempty" gorm:"type:text;column:campaign_name"`
// 	Description     string  `json:"description,omitempty" gorm:"type:text;column:description"`
// 	AboutProject    string  `json:"about_project,omitempty" gorm:"type:text;column:about_project"`
// 	Summary         string  `json:"summary,omitempty" gorm:"type:text;column:summary"`
// 	VenturePurpose  string  `json:"venture_purpose,omitempty" gorm:"type:text;column:venture_purpose"`
// 	VenturePhase    string  `json:"venture_phase,omitempty" gorm:"type:text;column:venture_phase"`
// 	CreatedAt       time.Time `json:"created_at,omitempty" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
// 	UpdatedAt       time.Time `json:"updated_at,omitempty" gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

// 	// User tablosu ile ilişkilendirme
// 	User User `json:"-" gorm:"foreignKey:UserId;references:ID"`
// }

// // TableName fonksiyonu tablo adını belirtir
// func (Campaign) TableName() string {
// 	return "campaigns"
// }
