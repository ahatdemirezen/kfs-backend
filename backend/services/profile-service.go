package services

import (
	"fmt"
	"kfs-backend/database"
	"kfs-backend/models"
	"mime/multipart" // ✅ Doğru paket eklendi
	"time"

	"github.com/gofiber/fiber/v2"
)

// Profil güncelleme servisi (Fotoğraf desteği eklendi)
func UpdateProfile(userId uint, fileHeader *multipart.FileHeader, website, identityNumber, birthDateStr, gender, academicTitle string) (*models.Profile, error) {
	var profile models.Profile

	// Kullanıcının profilini bul
	if err := database.DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Profil bulunamadı")
	}

	// Eğer dosya yüklenmişse, önceki dosyayı S3'ten sil ve yenisini yükle
	if fileHeader != nil {
		// Önceki dosya varsa sil
		if profile.PhotoURL != "" {
			s3Service, err := NewS3Service()
			if err == nil {
				_ = s3Service.DeleteFile(profile.PhotoURL)
			}
		}

		// Yeni dosyayı yükle
		photoURL, err := UploadProfilePicture(userId, fileHeader)
		if err != nil {
			return nil, err
		}
		profile.PhotoURL = photoURL // ✅ Yeni URL profilde güncellendi
	}

	// Doğum tarihini string'den time.Time formatına dönüştür
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

	// Güncellemeyi veritabanına kaydet
	if err := database.DB.Save(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Profil güncellenirken bir hata oluştu")
	}

	return &profile, nil
}

// **Sadece profil fotoğrafını güncelleyen fonksiyon**
func UpdateProfilePhoto(userId uint, fileHeader *multipart.FileHeader) (*models.Profile, error) {
	var profile models.Profile

	// Kullanıcının profilini bul
	if err := database.DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Profil bulunamadı")
	}

	// Eğer dosya yüklenmişse, önceki dosyayı S3'ten sil ve yenisini yükle
	if profile.PhotoURL != "" {
		s3Service, err := NewS3Service()
		if err == nil {
			_ = s3Service.DeleteFile(profile.PhotoURL)
		}
	}

	// Yeni dosyayı yükle
	photoURL, err := UploadProfilePicture(userId, fileHeader)
	if err != nil {
		return nil, err
	}

	// Güncellenen fotoğraf URL'sini kaydet
	profile.PhotoURL = photoURL

	// Güncellemeyi veritabanına kaydet
	if err := database.DB.Save(&profile).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Profil güncellenirken bir hata oluştu")
	}

	return &profile, nil
}

// **Profil fotoğrafını S3'e yükleyen fonksiyon**
func UploadProfilePicture(userId uint, fileHeader *multipart.FileHeader) (string, error) {
	// S3 servisini başlat
	s3Service, err := NewS3Service()
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	// Klasör yolunu belirle (örneğin: profiles/user_3/profile_picture.jpg)
	filePath := fmt.Sprintf("profiles/user_%d/profile_picture.jpg", userId)

	// Dosyayı S3'e yükle
	fileKey, err := s3Service.UploadFile(fileHeader, filePath)
	if err != nil {
		return "", fiber.NewError(fiber.StatusInternalServerError, "Dosya yükleme başarısız")
	}

	return fileKey, nil
}
