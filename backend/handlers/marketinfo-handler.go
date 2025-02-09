package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// MarketInfo oluştur
func CreateMarketInfo(c *fiber.Ctx) error {
	var marketInfo models.MarketInfo

	if err := c.BodyParser(&marketInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateMarketInfo(&marketInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(marketInfo)
}

// ID'ye göre MarketInfo getir
func GetMarketInfo(c *fiber.Ctx) error {
	marketInfoId := c.Params("marketInfoId")

	id, err := strconv.ParseUint(marketInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid marketInfoId",
		})
	}

	marketInfo, err := services.GetMarketInfoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(marketInfo)
}

// Belirli bir Campaign'e bağlı tüm MarketInfo kayıtlarını getir
func GetMarketInfosByCampaign(c *fiber.Ctx) error {
	campaignId := c.Params("campaignId")

	id, err := strconv.ParseUint(campaignId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid campaignId",
		})
	}

	marketInfos, err := services.GetMarketInfosByCampaignID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get market infos",
		})
	}

	return c.JSON(marketInfos)
}

// MarketInfo güncelle
func UpdateMarketInfo(c *fiber.Ctx) error {
	marketInfoId := c.Params("marketInfoId")

	id, err := strconv.ParseUint(marketInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid marketInfoId",
		})
	}

	marketInfo, err := services.GetMarketInfoByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&marketInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateMarketInfo(marketInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update market info",
		})
	}

	return c.JSON(marketInfo)
}

// MarketInfo sil
func DeleteMarketInfo(c *fiber.Ctx) error {
	marketInfoId := c.Params("marketInfoId")

	id, err := strconv.ParseUint(marketInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid marketInfoId",
		})
	}

	if err := services.DeleteMarketInfo(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete market info",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
