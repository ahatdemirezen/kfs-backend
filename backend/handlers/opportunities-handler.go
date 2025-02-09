package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Opportunity oluştur
func CreateOpportunity(c *fiber.Ctx) error {
	var opportunity models.Opportunity

	if err := c.BodyParser(&opportunity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateOpportunity(&opportunity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(opportunity)
}

// ID'ye göre Opportunity getir
func GetOpportunity(c *fiber.Ctx) error {
	opportunityId := c.Params("opportunityId")

	id, err := strconv.ParseUint(opportunityId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid opportunityId",
		})
	}

	opportunity, err := services.GetOpportunityByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(opportunity)
}

// Belirli bir AnalysisInfo'ya bağlı tüm Opportunities kayıtlarını getir
func GetOpportunitiesByAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	opportunities, err := services.GetOpportunitiesByAnalysisInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get opportunities",
		})
	}

	return c.JSON(opportunities)
}

// Opportunity güncelle
func UpdateOpportunity(c *fiber.Ctx) error {
	opportunityId := c.Params("opportunityId")

	id, err := strconv.ParseUint(opportunityId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid opportunityId",
		})
	}

	opportunity, err := services.GetOpportunityByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&opportunity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateOpportunity(opportunity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update opportunity",
		})
	}

	return c.JSON(opportunity)
}

// Opportunity sil
func DeleteOpportunity(c *fiber.Ctx) error {
	opportunityId := c.Params("opportunityId")

	id, err := strconv.ParseUint(opportunityId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid opportunityId",
		})
	}

	if err := services.DeleteOpportunity(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete opportunity",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Birden fazla Opportunity oluştur
func CreateMultipleOpportunities(c *fiber.Ctx) error {
	var opportunities []models.Opportunity

	if err := c.BodyParser(&opportunities); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateMultipleOpportunities(opportunities); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create opportunities",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Multiple Opportunities successfully created",
		"count":   len(opportunities),
	})
}
