package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupIncomeItemsRoutes(app *fiber.App) {
	// Initialize generic service for IncomeItems
	incomeItemsService := &services.GenericVentureService[models.IncomeItems]{}

	// Create route group
	incomeItemsGroup := app.Group("/api/income-items")

	// Create new IncomeItem
	incomeItemsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.IncomeItemsRequest
		return handlers.CreateVenture(c, incomeItemsService, req)
	})

	// Retrieve IncomeItem by ID
	incomeItemsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, incomeItemsService)
	})

	// List IncomeItems by CampaignId
	incomeItemsGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, incomeItemsService, "campaign_id")
	})

	// Update IncomeItem
	incomeItemsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.IncomeItemsRequest
		return handlers.UpdateVenture(c, incomeItemsService, req)
	})

	// Delete IncomeItem
	incomeItemsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, incomeItemsService)
	})
}
