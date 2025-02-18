package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupExplanationsRoutes(app *fiber.App) {
	// Initialize generic service for Explanations
	explanationsService := &services.GenericVentureService[models.Explanations]{}

	// Create route group
	explanationsGroup := app.Group("/api/explanations")

	// Create new Explanation
	explanationsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ExplanationsRequest
		return handlers.CreateVenture(c, explanationsService, req)
	})

	// Retrieve Explanation by ID
	explanationsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, explanationsService)
	})

	// List Explanations by CampaignId
	explanationsGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, explanationsService , "campaign_id")
	})

	// Update Explanation
	explanationsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ExplanationsRequest
		return handlers.UpdateVenture(c, explanationsService, req)
	})

	// Delete Explanation
	explanationsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, explanationsService)
	})
}
