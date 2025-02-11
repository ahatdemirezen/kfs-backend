package handlers

import (
	"kfs-backend/database"
	"kfs-backend/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateInvestmentRequest struct {
	Balance float64 `json:"balance"`
}

type InvestmentResponse struct {
	InvestmentId uint      `json:"InvestmentId"`
	UserId       uint      `json:"UserId"`
	CampaignId   uint      `json:"CampaignId"`
	Balance      float64   `json:"Balance"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

func CreateInvestment(c *fiber.Ctx) error {
	var request CreateInvestmentRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "İstek gövdesi ayrıştırılamadı",
		})
	}

	// User ID'yi Locals'tan al
	userID := c.Locals("userId").(uint)

	// Campaign ID'yi URL'den al
	campaignID, err := strconv.ParseUint(c.Params("campaignId"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz kampanya ID'si",
		})
	}

	// Yeni yatırım oluştur
	investment := models.Investment{
		UserId:     userID,
		CampaignId: uint(campaignID),
		Balance:    request.Balance,
	}

	// Veritabanına kaydet
	if err := database.DB.Create(&investment).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Yatırım oluşturulurken bir hata oluştu",
		})
	}

	response := InvestmentResponse{
		InvestmentId: investment.InvestmentId,
		UserId:       investment.UserId,
		CampaignId:   investment.CampaignId,
		Balance:      investment.Balance,
		CreatedAt:    investment.CreatedAt,
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Yatırım başarıyla oluşturuldu",
		"investment": response,
	})
}

func GetInvestments(c *fiber.Ctx) error {
	userID := c.Locals("userId").(uint)

	var investments []models.Investment
	if err := database.DB.Where("user_id = ?", userID).Find(&investments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Yatırımlar alınamadı",
		})
	}

	var response []InvestmentResponse
	for _, inv := range investments {
		response = append(response, InvestmentResponse{
			InvestmentId: inv.InvestmentId,
			UserId:       inv.UserId,
			CampaignId:   inv.CampaignId,
			Balance:      inv.Balance,
			CreatedAt:    inv.CreatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

func GetAllInvestments(c *fiber.Ctx) error {
	var investments []models.Investment
	if err := database.DB.Find(&investments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Yatırımlar alınamadı",
		})
	}

	var response []InvestmentResponse
	for _, inv := range investments {
		response = append(response, InvestmentResponse{
			InvestmentId: inv.InvestmentId,
			UserId:       inv.UserId,
			CampaignId:   inv.CampaignId,
			Balance:      inv.Balance,
			CreatedAt:    inv.CreatedAt,
		})
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
