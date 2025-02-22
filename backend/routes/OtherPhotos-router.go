package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupOtherPhotosRoutes(app *fiber.App) {
	// Initialize generic service for OtherPhotos
	otherPhotosService := &services.GenericVentureService[models.OtherPhotos]{}

	// Create route group
	otherPhotosGroup := app.Group("/api/other-photos")

	// Create new OtherPhoto
	otherPhotosGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.OtherPhotosRequest
		return handlers.CreateVenture(c, otherPhotosService, req)
	})

	// Retrieve OtherPhoto by ID
	otherPhotosGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, otherPhotosService)
	})

	// List OtherPhotos by VisualInfoId
	otherPhotosGroup.Get("/list/:visual_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, otherPhotosService, "visual_info_id")
	})

	// Update OtherPhoto
	otherPhotosGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.OtherPhotosRequest
		return handlers.UpdateVenture(c, otherPhotosService, req)
	})

	// Delete OtherPhoto
	otherPhotosGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, otherPhotosService)
	})
}
