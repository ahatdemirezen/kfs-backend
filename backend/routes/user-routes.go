package routes

import (
	"github.com/gofiber/fiber/v2"
	"kfs-backend/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	users := app.Group("/api/users")
	users.Put("/:userId", handlers.UpdateUser)
}

func SetupRegisterRoutes(app *fiber.App) {
	// Auth routes
	auth := app.Group("/api/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/send-verification-email", handlers.SendVerificationEmail)
} 