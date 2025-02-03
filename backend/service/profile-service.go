package service

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// Veritabanından profil bilgilerini alır
func GetProfileByUserId(userId string) (*models.Profile, error) {
	var profile models.Profile
	// user_id'ye göre profil bulur
	result := database.DB.Preload("User").Where("user_id = ?", userId).First(&profile)

	if result.Error != nil {
		return nil, errors.New("profil bulunamadı")
	}

	return &profile, nil
}
