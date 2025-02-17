package handlers

import (
	"kfs-backend/services"
	"math/rand"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type SendVerificationEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// Doğrulama kodlarını geçici olarak tutmak için map
var verificationCodes = make(map[string]string)
var codesLock sync.RWMutex

func SendVerificationEmail(c *fiber.Ctx) error {
	var req SendVerificationEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "İstek formatı geçersiz")
	}

	// 6 haneli rastgele bir doğrulama kodu oluştur
	verificationCode := strconv.Itoa(100000 + rand.Intn(900000))

	// Doğrulama kodunu geçici olarak sakla
	codesLock.Lock()
	verificationCodes[req.Email] = verificationCode
	codesLock.Unlock()

	// Email gönder
	if err := services.SendVerificationEmail(req.Email, verificationCode); err != nil {
		return err // Service'den gelen fiber.Error'u direkt olarak dön
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Doğrulama kodu başarıyla gönderildi",
		"success": true,
		"email":   req.Email,
	})
}

// Doğrulama kodunu kontrol etmek için yardımcı fonksiyon
func VerifyCode(email, code string) bool {
	codesLock.RLock()
	savedCode, exists := verificationCodes[email]
	codesLock.RUnlock()

	if exists && savedCode == code {
		codesLock.Lock()
		delete(verificationCodes, email) // Kullanılan kodu sil
		codesLock.Unlock()
		return true
	}
	return false
}
