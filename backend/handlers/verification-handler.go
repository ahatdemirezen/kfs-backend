package handlers

import (
	"kfs-backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

var request struct {
	IdentityNumber string `json:"identityNumber" validate:"required"`
	UserId         uint   `json:"userId" validate:"required"`
}

// UpdateUserVerificationStatus, kullanıcının doğrulama durumunu günceller
func UpdateUserVerificationStatus(c *fiber.Ctx) error {

	if err := c.BodyParser(&request); err != nil {
		log.Println("HATA: Body parse edilemedi -", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz istek formatı",
		})
	}

	if err := services.UpdateVerificationStatus(request.UserId, request.IdentityNumber); err != nil {
		log.Println("HATA: Kullanıcı doğrulama durumu güncellenemedi -", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Doğrulama durumu güncellenemedi",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Kullanıcı doğrulama durumu başarıyla güncellendi.",
	})
}
