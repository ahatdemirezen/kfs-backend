package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

// Profil rotalarını tanımlar
func SetupProfileRoutes(app *fiber.App) {
	profile := app.Group("/api/profile")

	// Profil bilgilerini getirme
	profile.Get("/:userId",
		middleware.AuthenticateMiddleware,
		middleware.IsUserMiddleware,
		handlers.GetProfileByUserId,
	)

	// Profil bilgilerini güncelleme
	profile.Put("/", handlers.UpdateProfile)
}
