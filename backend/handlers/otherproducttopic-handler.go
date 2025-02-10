package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// OtherProductTopic oluştur
func CreateOtherProductTopic(c *fiber.Ctx) error {
	var topic models.OtherProductTopic

	// İstek gövdesini parse et
	if err := c.BodyParser(&topic); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Servis fonksiyonunu çağır ve dönen topic'i al
	createdTopic, err := services.CreateOtherProductTopic(&topic)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Oluşturulan kaydı dön
	return c.Status(fiber.StatusCreated).JSON(createdTopic)
}

// ID'ye göre OtherProductTopic getir
func GetOtherProductTopic(c *fiber.Ctx) error {
	topicId := c.Params("topicId")

	id, err := strconv.ParseUint(topicId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid topicId",
		})
	}

	topic, err := services.GetOtherProductTopicByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(topic)
}

// Belirli bir ProductModelInfo'ya bağlı tüm OtherProductTopic kayıtlarını getir
func GetOtherProductTopicsByProductModelInfo(c *fiber.Ctx) error {
	productModelInfoId := c.Params("productModelInfoId")

	id, err := strconv.ParseUint(productModelInfoId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid productModelInfoId",
		})
	}

	topics, err := services.GetOtherProductTopicsByProductModelInfoID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get other product topics",
		})
	}

	return c.JSON(topics)
}

// OtherProductTopic güncelle
func UpdateOtherProductTopic(c *fiber.Ctx) error {
	topicId := c.Params("topicId")

	id, err := strconv.ParseUint(topicId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid topicId",
		})
	}

	topic, err := services.GetOtherProductTopicByID(uint(id))
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

	if err := services.UpdateOtherProductTopic(topic); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update topic",
		})
	}

	return c.JSON(topic)
}

// OtherProductTopic sil
func DeleteOtherProductTopic(c *fiber.Ctx) error {
	topicId := c.Params("topicId")

	id, err := strconv.ParseUint(topicId, 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid topicId",
		})
	}

	if err := services.DeleteOtherProductTopic(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete topic",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
