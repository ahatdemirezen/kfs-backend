package services

import (
	"kfs-backend/database"
	"kfs-backend/models"

	"github.com/gofiber/fiber/v2"
)

// CreateInvestment servisi
func CreateInvestment(userID uint, campaignID uint, balance float64) (*models.Investment, error) {
	// Kampanya var mı kontrolü
	var campaign models.Campaign
	if err := database.DB.First(&campaign, campaignID).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Belirtilen kampanya bulunamadı")
	}

	// Balance kontrolü
	if balance <= 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Yatırım miktarı 0'dan büyük olmalıdır")
	} 

	investment := models.Investment{
		UserId:     userID,
		CampaignId: campaignID,
		Balance:    balance,
	}

	if err := database.DB.Create(&investment).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Yatırım oluşturulurken bir hata oluştu")
	}

	return &investment, nil
}

// GetUserInvestments servisi
func GetUserInvestments(userID uint) ([]models.Investment, error) {
	var investments []models.Investment
	if err := database.DB.Where("user_id = ?", userID).Find(&investments).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Yatırımlar alınamadı")
	}

	return investments, nil
}

// GetAllInvestments servisi
func GetAllInvestments() ([]models.Investment, error) {
	var investments []models.Investment
	if err := database.DB.Find(&investments).Error; err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Tüm yatırımlar alınamadı")
	}

	return investments, nil
}
