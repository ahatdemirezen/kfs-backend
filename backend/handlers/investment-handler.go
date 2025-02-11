package handlers

import (
	"kfs-backend/services"
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
		return fiber.NewError(fiber.StatusBadRequest, "İstek gövdesi ayrıştırılamadı")
	}

	// User ID'yi Locals'tan al
	userID := c.Locals("userId").(uint)

	// Campaign ID'yi URL'den al
	campaignID, err := strconv.ParseUint(c.Params("campaignId"), 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Geçersiz kampanya ID'si")
	}

	// Servisi çağır
	investment, err := services.CreateInvestment(userID, uint(campaignID), request.Balance)
	if err != nil {
		return err // Service'den gelen hatayı doğrudan ilet
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

	// Servisi çağır
	investments, err := services.GetUserInvestments(userID)
	if err != nil {
		return err // Service'den gelen hatayı doğrudan ilet
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
	// Servisi çağır
	investments, err := services.GetAllInvestments()
	if err != nil {
		return err // Service'den gelen hatayı doğrudan ilet
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
