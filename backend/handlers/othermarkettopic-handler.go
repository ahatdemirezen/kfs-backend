package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// OtherMarketTopic oluştur
func CreateOtherMarketTopic(c *fiber.Ctx) error {
	var topic models.OtherMarketTopic

	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.CreateOtherMarketTopic(&topic); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(topic)
}

// ID'ye göre OtherMarketTopic getir
func GetOtherMarketTopic(c *fiber.Ctx) error {
	topicId := c.Params("topicId")

	id, err := strconv.ParseUint(topicId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid topicId",
		})
	}

	topic, err := services.GetOtherMarketTopicByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(topic)
}

// Belirli bir MarketInfo'ya bağlı tüm OtherMarketTopic kayıtlarını getir
func GetOtherMarketTopicsByMarketInfo(c *fiber.Ctx) error {
	marketInfoId := c.Params("marketInfoId")

	id, err := strconv.ParseUint(marketInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid marketInfoId",
		})
	}

	topics, err := services.GetOtherMarketTopicsByMarketInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get other market topics",
		})
	}

	return c.JSON(topics)
}

// OtherMarketTopic güncelle
func UpdateOtherMarketTopic(c *fiber.Ctx) error {
	topicId := c.Params("topicId")

	id, err := strconv.ParseUint(topicId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid topicId",
		})
	}

	topic, err := services.GetOtherMarketTopicByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := services.UpdateOtherMarketTopic(topic); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update topic",
		})
	}

	return c.JSON(topic)
}

// OtherMarketTopic sil
func DeleteOtherMarketTopic(c *fiber.Ctx) error {
	topicId := c.Params("topicId")

	id, err := strconv.ParseUint(topicId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid topicId",
		})
	}

	if err := services.DeleteOtherMarketTopic(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete topic",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
