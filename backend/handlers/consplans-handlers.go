package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ConsPlan oluştur
func CreateConsPlan(c *fiber.Ctx) error {
	var consPlan models.ConsPlan

	if err := c.BodyParser(&consPlan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateConsPlan(&consPlan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(consPlan)
}

// ID'ye göre ConsPlan getir
func GetConsPlan(c *fiber.Ctx) error {
	consPlanId := c.Params("consPlanId")

	id, err := strconv.ParseUint(consPlanId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid consPlanId",
		})
	}

	consPlan, err := services.GetConsPlanByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(consPlan)
}

// Belirli bir AnalysisInfo'ya bağlı tüm ConsPlan kayıtlarını getir
func GetConsPlansByAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	consPlans, err := services.GetConsPlansByAnalysisInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get cons plans",
		})
	}

	return c.JSON(consPlans)
}

// ConsPlan güncelle
func UpdateConsPlan(c *fiber.Ctx) error {
	consPlanId := c.Params("consPlanId")

	id, err := strconv.ParseUint(consPlanId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid consPlanId",
		})
	}

	consPlan, err := services.GetConsPlanByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&consPlan); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateConsPlan(consPlan); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update cons plan",
		})
	}

	return c.JSON(consPlan)
}

// ConsPlan sil
func DeleteConsPlan(c *fiber.Ctx) error {
	consPlanId := c.Params("consPlanId")

	id, err := strconv.ParseUint(consPlanId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid consPlanId",
		})
	}

	if err := services.DeleteConsPlan(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete cons plan",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Birden fazla ConsPlan oluştur
func CreateMultipleConsPlans(c *fiber.Ctx) error {
	var consPlans []models.ConsPlan

	if err := c.BodyParser(&consPlans); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateMultipleConsPlans(consPlans); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create cons plans",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Multiple ConsPlans successfully created",
		"count":   len(consPlans),
	})
}
