package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

// FinancialCategory Routes
func SetupFinancialCategoryRoutes(app *fiber.App) {
	categoryService := &services.GenericVentureService[models.FinancialCategory]{}
	categoryGroup := app.Group("/api/financial-categories")

	categoryGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.FinancialCategoryRequest
		return handlers.CreateVenture(c, categoryService, req)
	})

	categoryGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, categoryService)
	})

	categoryGroup.Get("/list", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, categoryService, "category_id")
	})

	categoryGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.FinancialCategoryRequest
		return handlers.UpdateVenture(c, categoryService, req)
	})

	categoryGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, categoryService)
	})
}
