package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Pros oluştur
func CreatePro(c *fiber.Ctx) error {
	var pro models.Pros

	// İstek verisini parse et
	if err := c.BodyParser(&pro); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Service fonksiyonunu çağır
	if err := services.CreatePro(&pro); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(pro)
}

// Pros oku
func GetPro(c *fiber.Ctx) error {
	proId := c.Params("proId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(proId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid proId",
		})
	}

	// Service fonksiyonunu çağır
	pro, err := services.GetProByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(pro)
}

// Belirli bir AnalysisInfo'ya bağlı tüm Pros kayıtlarını getir
func GetProsByAnalysisInfo(c *fiber.Ctx) error {
	analysisInfoId := c.Params("analysisInfoId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(analysisInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid analysisInfoId",
		})
	}

	// Service fonksiyonunu çağır
	pros, err := services.GetProsByAnalysisInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get pros",
		})
	}

	return c.JSON(pros)
}

// Pros güncelle
func UpdatePro(c *fiber.Ctx) error {
	proId := c.Params("proId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(proId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid proId",
		})
	}

	// Service katmanından mevcut veriyi getir
	pro, err := services.GetProByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Gelen isteği mevcut veriyle güncelle
	if err := c.BodyParser(&pro); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Güncellenmiş veriyi kaydet
	if err := services.UpdatePro(pro); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update pro",
		})
	}

	return c.JSON(pro)
}

// Pros sil
func DeletePro(c *fiber.Ctx) error {
	proId := c.Params("proId")

	// ID'yi uint'e çevir
	id, err := strconv.ParseUint(proId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid proId",
		})
	}

	// Service katmanını kullanarak silme işlemi yap
	if err := services.DeletePro(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete pro",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Birden fazla Pros oluştur
func CreateMultiplePros(c *fiber.Ctx) error {
	var pros []models.Pros

	// İstek gövdesini parse et
	if err := c.BodyParser(&pros); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Service fonksiyonunu çağır
	if err := services.CreateMultiplePros(pros); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create pros",
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Multiple Pros successfully created",
		"count":   len(pros),
	})
}
