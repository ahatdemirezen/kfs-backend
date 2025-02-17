package handlers

import (
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Tüm kullanıcıları listeleme
func GetAllUsers(c *fiber.Ctx) error {
	users, err := services.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kullanıcılar listelenemedi",
		})
	}
	return c.JSON(users)
}

// Kullanıcıyı silme
func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz kullanıcı ID",
		})
	}

	err = services.DeleteUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kullanıcı silinemedi",
		})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı başarıyla silindi"})
}

// Kullanıcı rolünü güncelleme
func UpdateUserRole(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz kullanıcı ID",
		})
	}

	var request struct {
		Role string `json:"role"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz giriş verisi",
		})
	}

	err = services.UpdateUserRole(uint(id), request.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kullanıcı rolü güncellenemedi",
		})
	}

	return c.JSON(fiber.Map{"message": "Kullanıcı rolü başarıyla güncellendi"})
}
