package models

import (
	"time"
)

// Verification modeli
type Verification struct {
	VerificationId       uint      `gorm:"primaryKey;autoIncrement" json:"verificationId"`
	UserId              uint      `gorm:"not null;column:user_id" json:"userId"`
	Email               string    `gorm:"unique" json:"email"`
	Code                string    `json:"code"`
	IsPhoneVerified     bool      `gorm:"type:boolean;default:false;column:is_phone_verified" json:"isPhoneVerified"`
	IsEmailVerified     bool      `gorm:"type:boolean;default:false;column:is_email_verified" json:"isEmailVerified"`
	IsUserVerified      bool      `gorm:"type:boolean;default:false;column:is_user_verified" json:"isUserVerified"`
	IsLawApproved       bool      `gorm:"type:boolean;default:false;column:is_law_approved" json:"isLawApproved"`
	EmailVerificationCode string    `gorm:"column:email_verification_code"`
	EmailCodeExpiry      int64     `gorm:"column:email_code_expiry"`
	CreatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt           time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// User tablosu ile ilişkilendirme
	User User `gorm:"foreignKey:UserId;references:UserId"`
}

// TableName tablosunun adını belirtir
func (Verification) TableName() string {
	return "verifications"
}
