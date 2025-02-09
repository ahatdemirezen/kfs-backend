package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ProductModelInfo oluştur
func CreateProductModelInfo(c *fiber.Ctx) error {
	var info models.ProductModelInfo

	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Service fonksiyonunu çağır
	if err := services.CreateProductModelInfo(&info); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// ProductModelInfo ve ilişkili Campaign bilgilerini döndür
	return c.Status(fiber.StatusCreated).JSON(info)
}

// ID'ye göre ProductModelInfo getir
func GetProductModelInfo(c *fiber.Ctx) error {
	infoId := c.Params("infoId")

	id, err := strconv.ParseUint(infoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid infoId",
		})
	}

	info, err := services.GetProductModelInfoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(info)
}

// Belirli bir Campaign'e bağlı tüm ProductModelInfo kayıtlarını getir
func GetProductModelInfosByCampaign(c *fiber.Ctx) error {
	campaignId := c.Params("campaignId")

	id, err := strconv.ParseUint(campaignId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaignId",
		})
	}

	infos, err := services.GetProductModelInfosByCampaignID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get product model infos",
		})
	}

	return c.JSON(infos)
}

// ProductModelInfo güncelle
func UpdateProductModelInfo(c *fiber.Ctx) error {
	infoId := c.Params("infoId")

	id, err := strconv.ParseUint(infoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid infoId",
		})
	}

	info, err := services.GetProductModelInfoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&info); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateProductModelInfo(info); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update product model info",
		})
	}

	return c.JSON(info)
}

// ProductModelInfo sil
func DeleteProductModelInfo(c *fiber.Ctx) error {
	infoId := c.Params("infoId")

	id, err := strconv.ParseUint(infoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid infoId",
		})
	}

	if err := services.DeleteProductModelInfo(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete product model info",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
