package models

import (
	"time"
)

// Roles modeli
type Role struct {
	RoleId    uint      `gorm:"primaryKey;autoIncrement;column:id"`
	UserId    uint      `gorm:"not null;column:user_id"` // Foreign key
	Role      string    `gorm:"type:text;not null;column:role"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:updated_at"`

	// User tablosu ile ilişkilendirme
	User User `gorm:"foreignKey:UserId;references:UserId"`
}

// TableName tablosunun adını belirtir
func (Role) TableName() string {
	return "roles"
}
