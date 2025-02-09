package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Threat oluştur
func CreateThreat(c *fiber.Ctx) error {
	var threat models.Threat

	if err := c.BodyParser(&threat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateThreat(&threat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(threat)
}

// ID'ye göre Threat getir
func GetThreat(c *fiber.Ctx) error {
	threatId := c.Params("threatId")

	id, err := strconv.ParseUint(threatId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threatId",
		})
	}

	threat, err := services.GetThreatByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(threat)
}

// Belirli bir AnalysisInfo'ya bağlı tüm Threat kayıtlarını getir
func GetThreatsByAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	threats, err := services.GetThreatsByAnalysisInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get threats",
		})
	}

	return c.JSON(threats)
}

// Threat güncelle
func UpdateThreat(c *fiber.Ctx) error {
	threatId := c.Params("threatId")

	id, err := strconv.ParseUint(threatId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threatId",
		})
	}

	threat, err := services.GetThreatByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&threat); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateThreat(threat); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update threat",
		})
	}

	return c.JSON(threat)
}

// Threat sil
func DeleteThreat(c *fiber.Ctx) error {
	threatId := c.Params("threatId")

	id, err := strconv.ParseUint(threatId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threatId",
		})
	}

	if err := services.DeleteThreat(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete threat",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Birden fazla Threat oluştur
func CreateMultipleThreats(c *fiber.Ctx) error {
	var threats []models.Threat

	if err := c.BodyParser(&threats); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateMultipleThreats(threats); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create threats",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Multiple Threats successfully created",
		"count":   len(threats),
	})
}
