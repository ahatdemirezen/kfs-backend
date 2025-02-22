package services

import (
	"kfs-backend/database"
	"kfs-backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Profil güncelleme servisi
func UpdateProfile(userId uint, photoURL, website, identityNumber, birthDateStr, gender, academicTitle string) (*models.Profile, error) {
	var profile models.Profile

	// Kullanıcının profilini bul
	if err := database.DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Profil bulunamadı")
	}

	// Doğum tarihini string'den time.Time formatına dönüştür
	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Doğum tarihi formatı geçersiz. Beklenen format: YYYY-MM-DD")
	}

	// Profili güncelle
	profile.PhotoURL = photoURL
	profile.Website = website
	profile.IdentityNumber = identityNumber
	profile.BirthDate = birthDate
	profile.Gender = gender
	profile.AcademicTitle = academicTitle

	// Güncellemeyi veritabanına kaydet
	if err := database.DB.Save(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Profil güncellenirken bir hata oluştu")
	}

	return &profile, nil
}
