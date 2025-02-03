package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupAuthRoutes -> /login, /logout
func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/api/auth")

	auth.Post("/login", handlers.Login)
	auth.Post("/logout", handlers.Logout)
}
