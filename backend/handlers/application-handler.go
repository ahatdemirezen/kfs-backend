package handlers

import (
	"encoding/json"
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RoleAppHandler(c *fiber.Ctx) error {
	// UserID'yi Header'dan al
	userIdStr := c.Get("UserID")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz UserID")
	}

	// Kullanıcı doğrulamasını kontrol et
	isVerified, err := services.IsUserVerified(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Doğrulama kontrol hatası")
	}
	if !isVerified {
		return c.Status(fiber.StatusUnauthorized).SendString("Kullanıcı doğrulanmamış")
	}

	// Request body'den applicationType al
	var requestBody struct {
		ApplicationType string `json:"applicationType"`
	}
	err = json.Unmarshal(c.Body(), &requestBody)
	if err != nil || (requestBody.ApplicationType != "entrepreneur" && requestBody.ApplicationType != "investor") {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz applicationType")
	}

	// Role başvurusunu oluştur
	err = services.CreateRoleApplication(userId, requestBody.ApplicationType)
	if err != nil {
		if err.Error() == "Kullanıcı bu tür için zaten başvuru yapmış" {
			return c.Status(fiber.StatusConflict).SendString(err.Error()) // 409 Conflict
		}
		return c.Status(fiber.StatusInternalServerError).SendString("Role başvurusu oluşturulamadı")
	}

	return c.Status(fiber.StatusCreated).SendString("Role başvurusu başarıyla oluşturuldu")
}
