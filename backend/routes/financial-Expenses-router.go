package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

// FinancialExpense Routes
func SetupFinancialExpenseRoutes(app *fiber.App) {
	expenseService := &services.GenericVentureService[models.FinancialExpense]{}
	expenseGroup := app.Group("/api/financial-expenses")

	expenseGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.FinancialExpenseRequest
		return handlers.CreateVenture(c, expenseService, req)
	})

	expenseGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, expenseService)
	})

	expenseGroup.Get("/list/:campaign_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, expenseService, "campaign_id")
	})

	expenseGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.FinancialExpenseRequest
		return handlers.UpdateVenture(c, expenseService, req)
	})

	expenseGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, expenseService)
	})
}