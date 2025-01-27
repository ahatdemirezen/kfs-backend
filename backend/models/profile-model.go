package models

import (
	"time"
)

// Profile kullanıcı profil bilgileri
type Profile struct {
	ProfileId      uint      `gorm:"primaryKey;autoIncrement;column:id"`
	UserId         uint      `gorm:"not null;uniqueIndex;column:user_id"`
	PhotoURL       string    `gorm:"type:varchar(255);column:photo_url"`
	Website        string    `gorm:"type:varchar(255);column:website"`
	IdentityNumber string    `gorm:"type:varchar(11);column:identity_number"` // TC Kimlik
	BirthDate      time.Time `gorm:"type:date;column:birth_date"`
	Gender         string    `gorm:"type:varchar(20);column:gender"`
	AcademicTitle  string    `gorm:"type:varchar(100);column:academic_title"`
	CreatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt      time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`
	User           User      `gorm:"foreignKey:UserId;references:UserId"`
}

// TableName tablosunun adını belirtir
func (Profile) TableName() string {
	return "profiles"
}
