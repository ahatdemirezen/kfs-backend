package models

import (
	"time"
)

// Verification modeli
type Verification struct {
	VerificationId  uint      `gorm:"primaryKey;autoIncrement;column:id"` // Primary key
	UserId          uint      `gorm:"not null;column:user_id"`            // Foreign key
	IsPhoneVerified bool      `gorm:"type:boolean;default:false;column:is_phone_verified"`
	IsEmailVerified bool      `gorm:"type:boolean;default:false;column:is_email_verified"`
	IsUserVerified  bool      `gorm:"type:boolean;default:false;column:is_user_verified"`
	IsLawApproved   bool      `gorm:"type:boolean;default:false;column:is_law_approved"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// User tablosu ile ilişkilendirme
	User User `gorm:"foreignKey:UserId;references:UserId"`
}

// TableName tablosunun adını belirtir
func (Verification) TableName() string {
	return "verifications"
}
