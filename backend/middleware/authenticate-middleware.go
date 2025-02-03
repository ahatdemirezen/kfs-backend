package middleware

import (
	"log"

	"kfs-backend/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AuthenticateMiddleware sadece JWT token doğrulaması yapar ve claims'leri locals'e ekler
func AuthenticateMiddleware(c *fiber.Ctx) error {
	// Cookie'den access_token al
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Erişim tokenı bulunamadı",
		})
	}

	// .env dosyasından secret key'i al
	secret := config.AppConfig.JwtSecret
	if secret == "" {
		log.Println("JWT_SECRET env değişkeni eksik")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Sunucu yapılandırma hatası",
		})
	}

	// Token'ı doğrula
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		// Token imza yöntemi kontrolü
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Geçersiz token imza yöntemi")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		log.Println("Token doğrulama hatası:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Geçersiz veya süresi dolmuş token",
		})
	}

	// JWT'deki claim'leri al
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token verileri alınamadı",
		})
	}
    // nodejs ' de req.user ' a eklenir go fiber ' de c.Locals' e eklenir.
	// Claims'leri Locals'e ekle
	c.Locals("claims", claims)

	// UserId'yi al ve ekle
	userId, ok := claims["userId"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token içerisinde userId bulunamadı",
		})
	}
	c.Locals("userId", uint(userId))

	// Role'ü al ve ekle
	role, ok := claims["role"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token içerisinde role bilgisi bulunamadı",
		})
	}
	c.Locals("role", role)

	log.Printf("Token doğrulandı - UserID: %d, Role: %s", uint(userId), role)

	return c.Next()
}
