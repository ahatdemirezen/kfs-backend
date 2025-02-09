package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Cons oluştur
func CreateCon(c *fiber.Ctx) error {
	var con models.Cons

	if err := c.BodyParser(&con); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateCon(&con); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(con)
}

// ID'ye göre Cons getir
func GetCon(c *fiber.Ctx) error {
	conId := c.Params("conId")

	id, err := strconv.ParseUint(conId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid conId",
		})
	}

	con, err := services.GetConByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(con)
}

// Belirli bir AnalysisInfo'ya bağlı tüm Cons kayıtlarını getir
func GetConsByAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	cons, err := services.GetConsByAnalysisInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get cons",
		})
	}

	return c.JSON(cons)
}

// Cons güncelle
func UpdateCon(c *fiber.Ctx) error {
	conId := c.Params("conId")

	id, err := strconv.ParseUint(conId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid conId",
		})
	}

	con, err := services.GetConByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&con); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateCon(con); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update con",
		})
	}

	return c.JSON(con)
}

// Cons sil
func DeleteCon(c *fiber.Ctx) error {
	conId := c.Params("conId")

	id, err := strconv.ParseUint(conId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid conId",
		})
	}

	if err := services.DeleteCon(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete con",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Birden fazla Cons oluştur
func CreateMultipleCons(c *fiber.Ctx) error {
	var cons []models.Cons

	if err := c.BodyParser(&cons); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateMultipleCons(cons); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create cons",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Multiple Cons successfully created",
		"count":   len(cons),
	})
}
