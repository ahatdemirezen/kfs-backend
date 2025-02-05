package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
	"time"
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

// Profil güncelleme servisi
func UpdateProfile(userId uint, photoURL, website, identityNumber, birthDateStr, gender, academicTitle string) (*models.Profile, error) {
	var profile models.Profile

	// Kullanıcının profilini bul
	if err := database.DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		return nil, errors.New("profil bulunamadı")
	}

	// Doğum tarihini string'den time.Time formatına dönüştür
	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		return nil, errors.New("doğum tarihi formatı geçersiz. beklenen format: YYYY-MM-DD")
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
		return nil, errors.New("profil güncellenirken bir hata oluştu")
	}

	return &profile, nil
}

// Yeni profil oluşturma servisi
func CreateProfile(userId uint, photoURL, website, identityNumber, birthDateStr, gender, academicTitle string) (*models.Profile, error) {
	// Kullanıcının zaten profili var mı kontrol et
	var existingProfile models.Profile
	if err := database.DB.Where("user_id = ?", userId).First(&existingProfile).Error; err == nil {
		return nil, errors.New("bu kullanıcı için zaten bir profil mevcut")
	}

	// Doğum tarihini string'den time.Time formatına dönüştür
	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		return nil, errors.New("doğum tarihi formatı geçersiz. beklenen format: YYYY-MM-DD")
	}

	// Yeni profil oluştur
	profile := models.Profile{
		UserId:         userId,
		PhotoURL:       photoURL,
		Website:        website,
		IdentityNumber: identityNumber,
		BirthDate:      birthDate,
		Gender:         gender,
		AcademicTitle:  academicTitle,
	}

	// Profili veritabanına kaydet
	if err := database.DB.Create(&profile).Error; err != nil {
		return nil, errors.New("profil oluşturulurken bir hata oluştu")
	}

	return &profile, nil
}
