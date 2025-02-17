package services

import (
	"fmt"
	"log"

	"kfs-backend/database"
	"kfs-backend/models"
)

// UpdateVerificationStatus, kullanıcının IsUserVerified durumunu günceller
func UpdateVerificationStatus(userId uint, identityNumber string) error {
	db := database.DB

	// Kullanıcının doğrulama kaydını bul
	var verification models.Verification
	if err := db.Where("user_id = ?", userId).Take(&verification).Error; err != nil {
		log.Println("HATA: Kullanıcının doğrulama kaydı bulunamadı -", err, "UserID:", userId)
		return err
	}
	

	//Eğer kimlik numarası daha önce doğrulanmışsa hata döndür
	if verification.IsUserVerified {
		log.Println("HATA: Kullanıcı zaten doğrulanmış -UserID", userId)
		return fmt.Errorf("Kullanıcı zaten doğrulanmış")
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
