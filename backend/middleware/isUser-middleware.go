package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// IsUserMiddleware, kullanıcı rolü kontrolü yapar
func IsUserMiddleware(c *fiber.Ctx) error {
	// Role listesini string array olarak al
	roles, ok := c.Locals("roles").([]string)
	if !ok {
		log.Println("HATA: Rol bilgisi bulunamadı veya yanlış formatta")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Rol bilgisi bulunamadı",
		})
	}

	// Kullanıcının rollerinde "bireysel" veya "kurumsal" olup olmadığını kontrol et
	isAuthorized := false
	for _, role := range roles {
		if role == "individual" || role == "corporate" {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		log.Println("HATA: Yetkisiz erişim - Kullanıcı rolü uygun değil")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Bu endpoint için yetkiniz bulunmamaktadır",
		})
	}

	log.Printf("Kullanıcı yetkilendirildi - Roller: %v", roles)

	return c.Next()
}
