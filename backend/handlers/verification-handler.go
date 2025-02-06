package handlers

import (

	"kfs-backend/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	IdentityNumber string `json:"identityNumber" validate:"required"`
	UserId         uint   `json:"userId" validate:"required"`
}

// UpdateUserVerificationStatus, kullanıcının doğrulama durumunu günceller
func UpdateUserVerificationStatus(c *fiber.Ctx) error {
    var request struct {
        IdentityNumber string `json:"identityNumber"`
        UserId         uint   `json:"userId"`
    }

	userId := c.Locals("userId").(uint)
    if err := c.BodyParser(&request); err != nil {
        log.Println("HATA: Body parse edilemedi -", err)
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Geçersiz istek formatı",
        })
    }

    log.Printf("DEBUG: Gelen UserID: %d, IdentityNumber: %s", userId, request.IdentityNumber)

    if userId == 0 {
        log.Println("HATA: Geçersiz userId, 0 geldi!")
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Geçersiz userId, 0 olamaz",
        })
    }

    if err := services.UpdateVerificationStatus(userId, request.IdentityNumber); err != nil {
        log.Println("HATA: Kullanıcı doğrulama durumu güncellenemedi -", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Doğrulama durumu güncellenemedi: " + err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "Kullanıcı doğrulama durumu başarıyla güncellendi.",
    })
}
