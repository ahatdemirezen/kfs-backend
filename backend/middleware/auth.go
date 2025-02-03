package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// AuthRequired -> cookie'den access_token okuyup kontrol eden middleware
func AuthRequired(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	if accessToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Access token cookie bulunamadı",
		})
	}

	// Token'ı parse et
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Geçersiz imza yöntemi: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token geçersiz veya parse edilemedi",
		})
	}

	// Exp süresi dolmuş mu?
	claims := token.Claims.(jwt.MapClaims)
	exp := int64(claims["exp"].(float64))
	if time.Now().Unix() > exp {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Token süresi dolmuş",
		})
	}

	// userId => c.Locals içine kaydediyoruz
	userIdFloat := claims["userId"].(float64)
	c.Locals("userId", uint(userIdFloat))

	return c.Next()
}
