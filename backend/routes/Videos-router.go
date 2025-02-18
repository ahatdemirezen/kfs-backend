package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupVideosRoutes(app *fiber.App) {
	// Initialize generic service for Videos
	videosService := &services.GenericVentureService[models.Videos]{}

	// Create route group
	videosGroup := app.Group("/api/videos")

	// Create new Video
	videosGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.VideosRequest
		return handlers.CreateVenture(c, videosService, req)
	})

	// Retrieve Video by ID
	videosGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, videosService)
	})

	// List Videos by VisualInfoId
	videosGroup.Get("/list/:visual_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, videosService, "visual_info_id")
	})

	// Update Video
	videosGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.VideosRequest
		return handlers.UpdateVenture(c, videosService, req)
	})

	// Delete Video
	videosGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, videosService)
	})
}
