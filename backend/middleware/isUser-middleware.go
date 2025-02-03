package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// UserAuthMiddleware, kullanıcı rolü kontrolü yapar
func IsUserMiddleware(c *fiber.Ctx) error {
	// Role değerini al
	role, ok := c.Locals("role").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Rol bilgisi bulunamadı",
		})
	}

	// Role kontrolü yap
	if role != "bireysel" && role != "kurumsal" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Bu endpoint için yetkiniz bulunmamaktadır",
		})
	}

	log.Printf("Kullanıcı yetkilendirildi - Role: %s", role)

	return c.Next()
}
