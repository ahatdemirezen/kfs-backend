package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
	"log"
)

// Check if a user's application exists
func IsApplicationExists(userId int, applicationType string) (bool, error) {
	var count int64
	result := database.DB.Model(&models.RoleApplicationForm{}).
		Where("user_id = ? AND application_type = ?", userId, applicationType).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func IsUserVerified(userId int) (bool, error) {
	var verification models.Verification
	result := database.DB.Where("user_id = ?", userId).First(&verification)

	if result.Error != nil {
		log.Printf("Verification check error: %v", result.Error)
		return false, result.Error
	}

	// Güncellenmiş veriyi terminalde görmek için log ekleyelim
	log.Printf("DB'den çekilen veri - UserID: %d, isUserVerified: %v", userId, verification.IsUserVerified)

	return verification.IsUserVerified, nil
}

// Create a role application
func CreateRoleApplication(userId int, applicationType string) error {
	log.Printf("Başvuru süreci başladı - UserID: %d, ApplicationType: %s", userId, applicationType)

	// 1️⃣ Kullanıcı daha önce başvuru yapmış mı kontrol et
	exists, err := IsApplicationExists(userId, applicationType)
	if err != nil {
		log.Printf("Başvuru kontrol hatası - UserID: %d, Hata: %v", userId, err)
		return err
	}
	if exists {
		log.Printf("Kullanıcı zaten başvuru yapmış - UserID: %d, ApplicationType: %s", userId, applicationType)
		return errors.New("Kullanıcı bu tür için zaten başvuru yapmış")
	}

	// 2️⃣ Yeni başvuruyu oluştur
	application := models.RoleApplicationForm{
		UserId:          uint(userId),
		ApplicationType: applicationType,
	}

	log.Println("Yeni başvuru ekleniyor...")

	result := database.DB.Create(&application)
	if result.Error != nil {
		log.Printf("Başvuru eklenirken hata oluştu - UserID: %d, ApplicationType: %s, Hata: %v", userId, applicationType, result.Error)
		return errors.New("Role başvurusu oluşturulamadı")
	}

	log.Printf("Role başvurusu başarıyla oluşturuldu! - UserID: %d, ApplicationType: %s", userId, applicationType)
	return nil
}
