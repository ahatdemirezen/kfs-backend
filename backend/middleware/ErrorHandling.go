package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorHandler, uygulamadaki hataları yakalayan ve uygun yanıtları döndüren bir middleware
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Varsayılan hata kodu
	code := fiber.StatusInternalServerError

	// Eğer hata bir Fiber hatası ise
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	// Hata yanıtını oluştur
	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    code,
			"message": err.Error(),
		},
	})
}
