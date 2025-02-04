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

	// Profil oluşturma
	profile.Post("/",
		middleware.AuthenticateMiddleware,
		middleware.IsUserMiddleware,
		handlers.CreateProfile,
	)

	// Profil bilgilerini güncelleme
	profile.Put("/",
		middleware.AuthenticateMiddleware,
		middleware.IsUserMiddleware,
		handlers.UpdateProfile)
}
