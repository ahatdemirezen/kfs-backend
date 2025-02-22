package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	users := app.Group("/api/users")

	// Kullanıcı bilgilerini getirme
	users.Get("/me",
		middleware.AuthenticateMiddleware,
		middleware.IsUserMiddleware,
		handlers.GetUser,
	)

	users.Put("/:userId", handlers.UpdateUser)
}

func SetupRegisterRoutes(app *fiber.App) {
	// Auth routes
	auth := app.Group("/api/auth")
	auth.Post("/register", handlers.Register)
	auth.Post("/send-verification-email", handlers.SendVerificationEmail)
}
