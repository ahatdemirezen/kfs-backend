package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

// FinancialSubCategory Routes
func SetupFinancialSubCategoryRoutes(app *fiber.App) {
	subCategoryService := &services.GenericVentureService[models.FinancialSubCategory]{}
	subCategoryGroup := app.Group("/api/financial-sub-categories")

	subCategoryGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.FinancialSubCategoryRequest
		return handlers.CreateVenture(c, subCategoryService, req)
	})

	subCategoryGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, subCategoryService)
	})

	subCategoryGroup.Get("/list/:category_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, subCategoryService, "category_id")
	})

	subCategoryGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.FinancialSubCategoryRequest
		return handlers.UpdateVenture(c, subCategoryService, req)
	})

	subCategoryGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, subCategoryService)
	})
}
