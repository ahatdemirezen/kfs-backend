package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupProfitForecastRoutes(app *fiber.App) {
	// Initialize generic service for ProfitForecast
	profitForecastService := &services.GenericVentureService[models.ProfitForecast]{}

	// Create route group
	profitForecastGroup := app.Group("/api/profit-forecast")

	// Create new ProfitForecast
	profitForecastGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ProfitForecastRequest
		return handlers.CreateVenture(c, profitForecastService, req)
	})

	// Retrieve ProfitForecast by ID
	profitForecastGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, profitForecastService)
	})

	// List ProfitForecasts by CampaignId
	profitForecastGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, profitForecastService, "campaign_id")
	})

	// Update ProfitForecast
	profitForecastGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ProfitForecastRequest
		return handlers.UpdateVenture(c, profitForecastService, req)
	})

	// Delete ProfitForecast
	profitForecastGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, profitForecastService)
	})
}
