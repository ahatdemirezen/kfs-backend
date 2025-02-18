package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupVisualInfoRoutes(app *fiber.App) {
	// Initialize generic service for VisualInfo
	visualInfoService := &services.GenericVentureService[models.VisualInfo]{}

	// Create route group
	visualInfoGroup := app.Group("/api/visual-info")

	// Create new VisualInfo
	visualInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.VisualInfoRequest
		return handlers.CreateVenture(c, visualInfoService, req)
	})

	// Retrieve VisualInfo by ID
	visualInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, visualInfoService)
	})

	// List VisualInfo by CampaignId
	visualInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, visualInfoService, "campaign_id")
	})

	// Update VisualInfo
	visualInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.VisualInfoRequest
		return handlers.UpdateVenture(c, visualInfoService, req)
	})

	// Delete VisualInfo
	visualInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, visualInfoService)
	})
}