package services

import (
	"log"

	"kfs-backend/database"
	"kfs-backend/models"
)

// UpdateVerificationStatus, kullanıcının IsUserVerified durumunu günceller
func UpdateVerificationStatus(userId uint, identityNumber string) error {
	db := database.DB

	// Kullanıcının doğrulama kaydını bul
	var verification models.Verification
	if err := db.Where("user_id = ?", userId).First(&verification).Error; err != nil {
		log.Println("HATA: Kullanıcının doğrulama kaydı bulunamadı -", err)
		return err
	}

	// Güncelleme işlemi
	verification.IsUserVerified = true
	if err := db.Save(&verification).Error; err != nil {
		log.Println("HATA: Kullanıcının doğrulama durumu güncellenemedi -", err)
		return err
	}

	var profile models.Profile
	if err := db.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		log.Println("HATA: Kullanıcının profili kaydı bulunamadı -", err)
		return err
	}

	profile.IdentityNumber = identityNumber
	if err := db.Save(&profile).Error; err != nil {
		log.Println("HATA: Kullanıcının kimlik bilgisi güncellenemedi -", err)
		return err
	}

	log.Printf("Başarılı: Kullanıcı doğrulama durumu güncellendi - UserID: %d, IsUserVerified: %v", userId, true)
	return nil
}
