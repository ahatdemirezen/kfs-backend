package routes

import (
	"github.com/gofiber/fiber/v2"
	"kfs-backend/handlers"
)
// send
func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")

	// Public routes
	auth.Post("/register", handlers.Register)
	auth.Post("/verify-email", handlers.VerifyEmail)
} 