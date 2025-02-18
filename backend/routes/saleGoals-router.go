package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupSaleGoalsRoutes(app *fiber.App) {
	// Initialize generic service for SaleGoals
	saleGoalsService := &services.GenericVentureService[models.SaleGoals]{}

	// Create route group
	saleGoalsGroup := app.Group("/api/sale-goals")

	// Create new SaleGoal
	saleGoalsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.SaleGoalsRequest
		return handlers.CreateVenture(c, saleGoalsService, req)
	})

	// Retrieve SaleGoal by ID
	saleGoalsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, saleGoalsService)
	})

	// List SaleGoals by CampaignId
	saleGoalsGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, saleGoalsService, "campaign_id")
	})

	// Update SaleGoal
	saleGoalsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.SaleGoalsRequest
		return handlers.UpdateVenture(c, saleGoalsService, req)
	})

	// Delete SaleGoal
	saleGoalsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, saleGoalsService)
	})
}
