package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ThreatPlan oluştur
func CreateThreatPlan(c *fiber.Ctx) error {
	var threatPlan models.ThreatPlan

	if err := c.BodyParser(&threatPlan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateThreatPlan(&threatPlan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(threatPlan)
}

// ID'ye göre ThreatPlan getir
func GetThreatPlan(c *fiber.Ctx) error {
	threatPlanId := c.Params("threatPlanId")

	id, err := strconv.ParseUint(threatPlanId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threatPlanId",
		})
	}

	threatPlan, err := services.GetThreatPlanByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(threatPlan)
}

// Belirli bir AnalysisInfo'ya bağlı tüm ThreatPlan kayıtlarını getir
func GetThreatPlansByAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	threatPlans, err := services.GetThreatPlansByAnalysisInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get threat plans",
		})
	}

	return c.JSON(threatPlans)
}

// ThreatPlan güncelle
func UpdateThreatPlan(c *fiber.Ctx) error {
	threatPlanId := c.Params("threatPlanId")

	id, err := strconv.ParseUint(threatPlanId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threatPlanId",
		})
	}

	threatPlan, err := services.GetThreatPlanByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&threatPlan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateThreatPlan(threatPlan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update threat plan",
		})
	}

	return c.JSON(threatPlan)
}

// ThreatPlan sil
func DeleteThreatPlan(c *fiber.Ctx) error {
	threatPlanId := c.Params("threatPlanId")

	id, err := strconv.ParseUint(threatPlanId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threatPlanId",
		})
	}

	if err := services.DeleteThreatPlan(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete threat plan",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Birden fazla ThreatPlan oluştur
func CreateMultipleThreatPlans(c *fiber.Ctx) error {
	var threatPlans []models.ThreatPlan

	if err := c.BodyParser(&threatPlans); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateMultipleThreatPlans(threatPlans); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create threat plans",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Multiple ThreatPlans successfully created",
		"count":   len(threatPlans),
	})
}
