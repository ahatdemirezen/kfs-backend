package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// Profil rotalarını tanımlar
func SetupProfileRoutes(app *fiber.App) {
	profile := app.Group("/api/profile")

	// Profil bilgilerini güncelleme
	profile.Put("/",
		middleware.AuthenticateMiddleware,
		middleware.IsUserMiddleware,
		handlers.UpdateProfile)

	profile.Post("/upload-photo",
		middleware.AuthenticateMiddleware,
		middleware.IsUserMiddleware,
		handlers.UpdateProfilePhoto)
}
