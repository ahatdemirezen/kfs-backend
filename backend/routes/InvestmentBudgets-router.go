package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupInvestmentBudgetsRoutes(app *fiber.App) {
	// Initialize generic service for InvestmentBudgets
	investmentBudgetsService := &services.GenericVentureService[models.InvestmentBudgets]{}

	// Create route group
	investmentBudgetsGroup := app.Group("/api/investment-budgets")

	// Create new InvestmentBudgets
	investmentBudgetsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.InvestmentBudgetsRequest
		return handlers.CreateVenture(c, investmentBudgetsService, req)
	})

	// Retrieve InvestmentBudgets by ID
	investmentBudgetsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, investmentBudgetsService)
	})

	// List InvestmentBudgets by CampaignId
	investmentBudgetsGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, investmentBudgetsService, "campaign_id")
	})

	// Update InvestmentBudgets
	investmentBudgetsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.InvestmentBudgetsRequest
		return handlers.UpdateVenture(c, investmentBudgetsService, req)
	})

	// Delete InvestmentBudgets
	investmentBudgetsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, investmentBudgetsService)
	})
}
