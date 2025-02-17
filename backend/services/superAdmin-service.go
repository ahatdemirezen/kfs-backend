package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

//tüm kullanıcıları getir

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result:=database.DB.Find(&users)

	if result.Error!= nil {
		return nil, result.Error
	}
	return users, nil
}

//kullanıcıyı sil

func DeleteUser(userId uint) error {

	result:=database.DB.Delete(&models.User{}, userId)

	if result.Error!= nil {
		return result.Error
	}
	if result.RowsAffected==0 {
		return errors.New("Kullanıcı bulunamadı")
	}	
	return nil
}


// Kullanıcı rolünü güncelle
func UpdateUserRole(userId uint, newRole string) error {
	// Kullanıcı olup olmadığını kontrol et
	var user models.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		return errors.New("Kullanıcı bulunamadı")
	}

	// Önce mevcut rolü sil (Eğer tek bir rol tutuyorsak)
	database.DB.Where("user_id = ?", userId).Delete(&models.Role{})

	// Yeni rolü kaydet
	role := models.Role{
		UserId: userId,
		Role:   newRole,
	}
	result := database.DB.Create(&role)
	if result.Error != nil {
		return result.Error
	}

	return nil
}