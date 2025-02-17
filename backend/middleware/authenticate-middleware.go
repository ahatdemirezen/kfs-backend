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
		log.Println("HATA: Erişim tokenı bulunamadı")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Erişim tokenı bulunamadı",
		})
	}

	// .env dosyasından secret key'i al
	secret := config.AppConfig.JwtSecret
	if secret == "" {
		log.Println("HATA: JWT_SECRET env değişkeni eksik")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Sunucu yapılandırma hatası",
		})
	}

	// Token'ı doğrula
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		// Token imza yöntemi kontrolü
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Println("HATA: Geçersiz token imza yöntemi")
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Geçersiz token imza yöntemi")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		log.Println("HATA: Token doğrulama hatası:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Geçersiz veya süresi dolmuş token",
		})
	}

	// JWT'deki claim'leri al
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("HATA: Token verileri alınamadı")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token verileri alınamadı",
		})
	}

	// Claims'leri Locals'e ekle
	c.Locals("claims", claims)

	// UserId'yi al ve ekle
	userId, ok := claims["userId"].(float64)
	if !ok {
		log.Println("HATA: Token içerisinde userId bulunamadı")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token içerisinde userId bulunamadı",
		})
	}
	c.Locals("userId", uint(userId))

	// ProfileId'yi al ve ekle
	profileId, ok := claims["profileId"].(float64)
	if !ok {
		log.Println("HATA: Token içerisinde profileId bulunamadı")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token içerisinde profileId bulunamadı",
		})
	}
	c.Locals("profileId", uint(profileId))

	// Rolleri al ve ekle
	roles, ok := claims["roles"].([]interface{})
	if !ok {
		log.Println("UYARI: Token içerisinde role bilgisi bulunamadı, boş array atanıyor")
		roles = []interface{}{} // Varsayılan olarak boş array
	}

	// []interface{} -> []string dönüşümü
	roleStrings := make([]string, len(roles))
	for i, r := range roles {
		if roleStr, ok := r.(string); ok {
			roleStrings[i] = roleStr
		} else {
			log.Println("UYARI: Role tipi string değil, atlanıyor")
		}
	}
	c.Locals("roles", roleStrings)

	log.Printf("Token doğrulandı - UserID: %d, Roller: %v", uint(userId), roleStrings)

	return c.Next()
}
