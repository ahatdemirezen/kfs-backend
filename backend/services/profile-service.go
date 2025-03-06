package services

import (
	"fmt"
	"mime/multipart"
	"time"

	"kfs-backend/database"
	"kfs-backend/models"

	"github.com/gofiber/fiber/v2"
)

// Profil güncelleme servisi (Fotoğraf yükleme dahil)
func UpdateProfile(userId uint, fileHeader *multipart.FileHeader, website, identityNumber, birthDateStr, gender, academicTitle string) (*models.Profile, error) {
	var profile models.Profile

	// Kullanıcının profilini bul
	if err := database.DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Profil bulunamadı")
	}

	// Eğer yeni fotoğraf dosyası yüklenmişse
	if fileHeader != nil {
		// Önce mevcut fotoğrafı sil (varsa)
		if profile.PhotoURL != "" {
			s3Service, err := NewS3Service()
			if err == nil {
				_ = s3Service.DeleteFile(profile.PhotoURL) // Hata olursa loglanabilir ama önemli değil
			}
		}

		// Yeni fotoğrafı S3'e yükle
		photoURL, err := uploadProfilePicture(userId, fileHeader)
		if err != nil {
			return nil, err
		}
		profile.PhotoURL = photoURL // Yeni fotoğraf URL'sini kaydet
	}

	// Doğum tarihi formatını kontrol et ve dönüştür
	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Doğum tarihi formatı geçersiz. Beklenen format: YYYY-MM-DD")
	}

	// Profili güncelle
	profile.Website = website
	profile.IdentityNumber = identityNumber
	profile.BirthDate = birthDate
	profile.Gender = gender
	profile.AcademicTitle = academicTitle

	// Veritabanında güncelle
	if err := database.DB.Save(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Profil güncellenirken bir hata oluştu")
	}

	return &profile, nil
}

// **Profil fotoğrafını S3'e yükleme fonksiyonu**
func uploadProfilePicture(userId uint, fileHeader *multipart.FileHeader) (string, error) {
	s3Service, err := NewS3Service()
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	filePath := fmt.Sprintf("profiles/user_%d/profile_picture.jpg", userId)

	fileKey, err := s3Service.UploadFile(fileHeader, filePath)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Fotoğraf yüklenirken bir hata oluştu")
	}

	return fileKey, nil
}
