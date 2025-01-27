package models

import (
	"time"
)

// User modeli
type User struct {
	UserId      uint      `gorm:"primaryKey;autoIncrement;column:id"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null;column:email"`
	Password    string    `gorm:"type:varchar(255);not null;column:password"`
	Salt        string    `gorm:"type:varchar(255);not null;column:salt"` // Salt alanı eklendi
	Country     string    `gorm:"type:varchar(100);column:country"`
	Phone       string    `gorm:"type:varchar(15);column:phone"`
	FirstName   string    `gorm:"type:varchar(100);column:first_name"`
	LastName    string    `gorm:"type:varchar(100);column:last_name"`
	CompanyName string    `gorm:"type:varchar(255);column:company_name"`
	TaxOffice   string    `gorm:"type:varchar(255);column:tax_office"`
	TaxNumber   string    `gorm:"type:varchar(50);column:tax_number"`
	CreatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`
}

// TableName tablosunun adını belirtir
func (User) TableName() string {
	return "users"
}
