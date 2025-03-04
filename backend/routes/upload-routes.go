package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupGeneralRoutes(app *fiber.App) {
	// Generic service
	genericService := &services.GenericVentureService[any]{}

	// Genel dosya yükleme ve veritabanı güncelleme
	app.Post("/api/upload-file/:id", func(c *fiber.Ctx) error {
		return handlers.UploadFileAndUpdate(c, genericService)
	})
}
