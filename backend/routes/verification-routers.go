package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"

	"kfs-backend/middleware"
)

// SetupVerificationRoutes doğrulama ile ilgili route'ları tanımlar
func SetupVerificationRoutes(app *fiber.App) {
	verificationGroup := app.Group("/api/verification")

	// Kullanıcının doğrulama durumunu güncelle
	verificationGroup.Put("/update", middleware.AuthenticateMiddleware, handlers.UpdateUserVerificationStatus)
}
