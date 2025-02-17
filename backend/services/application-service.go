package services

import (
	"kfs-backend/database"
	"errors"
)

// Kullanıcının başvurusunun varlığını kontrol et
func IsApplicationExists(userId int, applicationType string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM roleapplicationforms
		WHERE userId = $1 AND applicationType = $2
	`
	var count int
	err := database.DB.Raw(query, userId, applicationType).Scan(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Kullanıcı doğrulama durumu kontrolü
func IsUserVerified(userId int) (bool, error) {
	query := "SELECT isUserVerified FROM verification WHERE userId = $1"
	var isVerified bool
	err := database.DB.Raw(query, userId).Scan(&isVerified).Error
	if err != nil {
		return false, err
	}
	return isVerified, nil
}

// Role başvurusu oluşturma
func CreateRoleApplication(userId int, applicationType string) error {
	// Başvurunun zaten var olup olmadığını kontrol et
	exists, err := IsApplicationExists(userId, applicationType)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Kullanıcı bu tür için zaten başvuru yapmış")
	}
	// Yeni başvuru ekle
	query := `
		INSERT INTO roleapplicationforms (userId, applicationType)
		VALUES ($1, $2)
	`
	result := database.DB.Exec(query, userId, applicationType)
	if result.Error != nil {
		return result.Error
	}
	if err != nil {
		return err
	}
	return nil
}
