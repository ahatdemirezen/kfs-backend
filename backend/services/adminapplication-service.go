package services

import (
	"kfs-backend/database"
	"kfs-backend/models"
)

// Service interface tanımı
type RoleApplicationService interface {
	GetAllRoleApplications() ([]models.RoleApplicationForm, error)
	UpdateRoleApplicationStatus(applicationId uint, status string) error
}

// Service yapısı
type roleApplicationService struct{}

// Yeni servis örneği oluştur
func NewRoleApplicationService() RoleApplicationService {
	return &roleApplicationService{}
}

// Tüm rol başvurularını getir
func (s *roleApplicationService) GetAllRoleApplications() ([]models.RoleApplicationForm, error) {
	var applications []models.RoleApplicationForm
	if err := database.DB.Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

// Başvuru statüsünü güncelle
func (s *roleApplicationService) UpdateRoleApplicationStatus(applicationId uint, status string) error {
	tx := database.DB.Begin()
	if err := tx.Model(&models.RoleApplicationForm{}).Where("id = ?", applicationId).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	if status == "accepted" {
		var application models.RoleApplicationForm
		if err := tx.First(&application, applicationId).Error; err != nil {
			tx.Rollback()
			return err
		}

		role := models.Role{UserId: application.UserId, Role: application.ApplicationType}
		if err := tx.Create(&role).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}
