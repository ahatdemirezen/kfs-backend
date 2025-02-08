package handlers

import (
	"log"
	"time"

	"kfs-backend/config"

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
	secure := config.AppConfig.NodeEnv == "production"
	db := database.DB

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "İstek formatı geçersiz")
	}

	// 1) Kullanıcıyı email ile bul
	var user models.User
	result := db.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email veya şifre hatalı")
	}

	// 2) Kullanıcının rollerini getir
	var roles []string
	db.Table("roles").Where("user_id = ?", user.UserId).Pluck("role", &roles)
	if len(roles) == 0 {
		roles = append(roles, "user") // Varsayılan rol
	}

	// 3) Kullanıcının profilini getir
	var profile models.Profile
	profileResult := db.Where("user_id = ?", user.UserId).First(&profile)
	var profileId uint
	if profileResult.Error != nil {
		// Profil bulunamadıysa yeni profil oluştur
		newProfile := models.Profile{
			UserId: user.UserId,
		}
		if err := db.Create(&newProfile).Error; err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Profil oluşturulamadı")
		}
		profileId = newProfile.ProfileId
	} else {
		profileId = profile.ProfileId
	}

	log.Printf("Bulunan roller: %+v", roles) // Roller array olarak loglanıyor

	// 4) Şifreyi kontrol et
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password+user.Salt)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Email veya şifre hatalı")
	}

	log.Println("Login başarılı, userID:", user.UserId, "Roller:", roles)

	// 5) Access ve Refresh token oluştur
	accessToken, err := generateJWT(user.UserId, profileId, roles, 15*time.Minute, "access")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Access token oluşturulamadı")
	}
	refreshToken, err := generateJWT(user.UserId, profileId, roles, 24*7*time.Hour, "refresh")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Refresh token oluşturulamadı")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "strict", // frontend ve backend farklı serverlarda alınırsa canlıya burası değiştirilecek
		Path:     "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(24 * 7 * time.Hour),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "strict", // frontend ve backend farklı serverlarda alınırsa canlıya burası değiştirilecek
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Başarılı giriş. Tokenlar cookie olarak setlendi.",
	})
}

func Logout(c *fiber.Ctx) error {
	secure := config.AppConfig.NodeEnv == "production"
	// Aynı cookie isimleriyle, geçmiş bir expire vererek yok edebiliriz
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "strict", // frontend ve backend farklı serverlarda alınırsa canlıya burası değiştirilecek
		Path:     "/",
	})
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "strict", // frontend ve backend farklı serverlarda alınırsa canlıya burası değiştirilecek
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"message": "Çıkış yapıldı. Token cookie'leri silindi.",
	})
}

// RefreshToken: Refresh token'ı doğrulayıp yeni access token üretir
func RefreshToken(c *fiber.Ctx) error {
	secure := config.AppConfig.NodeEnv == "production"
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Refresh token bulunamadı")
	}

	// Refresh token'ı doğrula
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Geçersiz token imza metodu")
		}
		return []byte(config.AppConfig.JwtSecretRefresh), nil
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Geçersiz refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Geçersiz token claims")
	}

	// Token tipini kontrol et
	tokenType, ok := claims["type"].(string)
	if !ok || tokenType != "refresh" {
		return fiber.NewError(fiber.StatusUnauthorized, "Geçersiz token tipi")
	}

	// Kullanıcı bilgilerini al
	userId := uint(claims["userId"].(float64))
	profileId := uint(claims["profileId"].(float64))
	roles, ok := claims["roles"].([]interface{})
	if !ok {
		return fiber.NewError(fiber.StatusInternalServerError, "Roller alınamadı")
	}

	// Interface dizisini string dizisine dönüştür
	rolesStr := make([]string, len(roles))
	for i, role := range roles {
		rolesStr[i] = role.(string)
	}

	// Yeni access token oluştur
	newAccessToken, err := generateJWT(userId, profileId, rolesStr, 15*time.Minute, "access")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Yeni access token oluşturulamadı")
	}

	// Yeni access token'ı cookie olarak ayarla
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		Expires:  time.Now().Add(15 * time.Minute),
		HTTPOnly: true,
		Secure:   secure,
		SameSite: "strict", // frontend ve backend farklı serverlarda alınırsa canlıya burası değiştirilecek
		Path:     "/",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Access token yenilendi",
	})
}

// generateJWT: userId, profileId, role ve süre bilgisiyle JWT üretimi
func generateJWT(userId uint, profileId uint, roles []string, duration time.Duration, tokenType string) (string, error) {
	claims := jwt.MapClaims{
		"userId":    userId,
		"profileId": profileId,
		"roles":     roles,
		"type":      tokenType,
		"exp":       time.Now().Add(duration).Unix(),
		"iat":       time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := config.AppConfig.JwtSecret
	if tokenType == "refresh" {
		secret = config.AppConfig.JwtSecretRefresh
	}

	if secret == "" {
		log.Println("UYARI: JWT secret boş!")
		return "", fiber.NewError(fiber.StatusInternalServerError, "JWT secret bulunamadı")
	}

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("JWT imzalama hatası: %v", err)
		return "", err
	}

	log.Printf("JWT oluşturuldu. Token tipi: %s", tokenType)
	return signedToken, nil
}
