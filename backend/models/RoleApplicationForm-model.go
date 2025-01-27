package models

import (
	"time"
)

// RoleApplicationForm modeli
type RoleApplicationForm struct {
	ApplicationId   uint      `gorm:"primaryKey;autoIncrement;column:id"` // Primary key
	ApplicationType string    `gorm:"type:varchar(255);not null;column:application_type"`
	UserId          uint      `gorm:"not null;column:user_id"` // Foreign key
	City            string    `gorm:"type:text;column:city"`
	Town            string    `gorm:"type:text;column:town"`
	Nationality     string    `gorm:"type:text;column:nationality"`
	PostalCode      string    `gorm:"type:text;column:postal_code"`
	Address         string    `gorm:"type:text;column:address"`
	Education       string    `gorm:"type:text;column:education"`
	Profession      string    `gorm:"type:text;column:profession"`
	Job             string    `gorm:"type:text;column:job"`
	Sector          string    `gorm:"type:text;column:sector"`
	Experience      string    `gorm:"type:text;column:experience"`
	Website         string    `gorm:"type:text;column:website"`
	Bank            string    `gorm:"type:text;column:bank"`
	IBAN            string    `gorm:"type:text;column:iban"`
	ResumeKey       string    `gorm:"type:text;column:resume_key"`
	MkkNumber       string    `gorm:"type:text;column:mkk_number"`
	IncomeStatement int       `gorm:"type:int;column:income_statement"`
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// User tablosu ile ilişkilendirme
	User User `gorm:"foreignKey:UserId;references:UserId"`
}

// TableName tablosunun adını belirtir
func (RoleApplicationForm) TableName() string {
	return "role_application_forms"
}
