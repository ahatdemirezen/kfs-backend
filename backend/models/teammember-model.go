package models

import (
	"time"
)

// TeamMember modeli
type TeamMember struct {
	TeamMemberId  uint      `gorm:"primaryKey;autoIncrement;column:team_member_id"`
	CampaignId    uint      `gorm:"not null;column:campaign_id"` // Foreign key
	Name          string    `gorm:"type:text;column:name"`
	Surname       string    `gorm:"type:text;column:surname"`
	Position      string    `gorm:"type:text;column:position"`
	ResumeKey     string    `gorm:"type:text;column:resume_key"`
	PhotoKey      string    `gorm:"type:text;column:photo_key"`
	Biography     string    `gorm:"type:text;column:biography"`
	Responsibility string   `gorm:"type:text;column:responsibility"`
	Profession    string    `gorm:"type:text;column:profession"`
	Relation      string    `gorm:"type:text;column:relation"`
	Email         string    `gorm:"type:text;column:email"`
	Instagram     string    `gorm:"type:text;column:instagram"`
	Twitter       string    `gorm:"type:text;column:twitter"`
	Linkedin      string    `gorm:"type:text;column:linkedin"`
	CreatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// Campaign tablosu ile ilişkilendirme
	Campaign Campaign `gorm:"foreignKey:CampaignId;references:CampaignId"`
}

// TableName tablosunun adını belirtir
func (TeamMember) TableName() string {
	return "team_members"
}
