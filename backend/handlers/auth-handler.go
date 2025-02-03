package handlers

import (
	"os"
	"time"
    "log"               // <-- Bunu ekleyin

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"kfs-backend/database"
	"kfs-backend/models"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Login fonksiyonu
func Login(c *fiber.Ctx) error {
	db := database.DB

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Body parse edilemedi",
		})
	}

// 1) Kullanıcıyı email ile bul
var user models.User
result := db.Where("email = ?", req.Email).First(&user)
if result.Error != nil {
    log.Println("DB'den kullanıcı çekilemedi. Email:", req.Email)
    log.Println("Hata içeriği:", result.Error)
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Email veya şifre hatalı (kullanıcı yok)",
    })
}

// 2) Şifreyi kontrol et
if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password+user.Salt)); err != nil {
    log.Println("Şifre kontrolü başarısız, userID:", user.UserId, "Hata:", err)
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Email veya şifre hatalı (şifre uyumsuz)",
    })
}

log.Println("Login başarılı, userID:", user.UserId)


	// 3) Access ve Refresh token oluştur
	accessToken, err := generateJWT(user.UserId, 15*time.Minute, "access")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Access token oluşturulamadı",
		})
	}
	refreshToken, err := generateJWT(user.UserId, 24*7*time.Hour, "refresh") // 7 gün
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Refresh token oluşturulamadı",
		})
	}

	// 4) Cookie'lere yaz
	// Access token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   false, // Prod'da true + HTTPS
		SameSite: "strict",
		Path:     "/",
	})

	// Refresh token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "strict",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "Başarılı giriş. Tokenlar cookie olarak setlendi.",
	})
}

// Logout fonksiyonu
func Logout(c *fiber.Ctx) error {
	// Aynı cookie isimleriyle, geçmiş bir expire vererek yok edebiliriz
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "strict",
		Path:     "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "strict",
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "Çıkış yapıldı. Token cookie'leri silindi.",
	})
}

// generateJWT: userId ve süre bilgisiyle basit JWT üretimi
func generateJWT(userId uint, duration time.Duration, tokenType string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"type":   tokenType,
		"exp":    time.Now().Add(duration).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// .env dosyasında tanımlı secret'ları çekelim
	secret := os.Getenv("JWT_SECRET")
	if tokenType == "refresh" {
		rSecret := os.Getenv("JWT_SECRET_REFRESH")
		if rSecret != "" {
			secret = rSecret
		}
	}

	return token.SignedString([]byte(secret))
}
