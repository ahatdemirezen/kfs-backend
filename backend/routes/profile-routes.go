package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

// Profil rotalarını tanımlar
func SetupProfileRoutes(app *fiber.App) {
	profile := app.Group("/api/profile")

	// Profil bilgilerini getirme
	profile.Get("/:userId", handlers.GetProfileByUserId)

	// Profil bilgilerini güncelleme
	profile.Put("/", handlers.UpdateProfile)
}
