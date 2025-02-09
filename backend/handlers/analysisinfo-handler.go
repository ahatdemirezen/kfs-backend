package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// AnalysisInfo oluştur
func CreateAnalysisInfo(c *fiber.Ctx) error {
	var analysisInfo models.AnalysisInfo

	// İstek verisini parse et
	if err := c.BodyParser(&analysisInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Service fonksiyonunu çağır
	if err := services.CreateAnalysisInfo(&analysisInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create analysis info",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(analysisInfo)
}

// AnalysisInfo oku
func GetAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	// Service fonksiyonunu çağır
	analysisInfo, err := services.GetAnalysisInfoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(analysisInfo)
}

// AnalysisInfo güncelle
func UpdateAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	// Service katmanından mevcut veriyi getir
	analysisInfo, err := services.GetAnalysisInfoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Gelen isteği mevcut veriyle güncelle
	if err := c.BodyParser(&analysisInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Güncellenmiş veriyi kaydet
	if err := services.UpdateAnalysisInfo(analysisInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update analysis info",
		})
	}

	return c.JSON(analysisInfo)
}

// AnalysisInfo sil
func DeleteAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	// Service katmanını kullanarak silme işlemi yap
	if err := services.DeleteAnalysisInfo(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete analysis info",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
