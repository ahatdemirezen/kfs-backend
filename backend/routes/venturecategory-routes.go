package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupVentureCategoryRoutes(app *fiber.App) {
	// Generic service'i VentureCategory için başlat
	categoryService := &services.GenericVentureService[models.VentureCategory]{}

	// Route gruplarını oluştur
	categoryGroup := app.Group("api/venture-categories")

	// Yeni VentureCategory oluştur
	categoryGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.VentureCategoryRequest
		return handlers.CreateVenture(c, categoryService, req)
	})

	// ID ile VentureCategory getir
	categoryGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, categoryService)
	})

	// Campaign ID'ye göre VentureCategory listele
	categoryGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, categoryService, "campaign_id")
	})

	// VentureCategory güncelle
	categoryGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.VentureCategoryRequest
		return handlers.UpdateVenture(c, categoryService, req)
	})

	// VentureCategory sil
	categoryGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, categoryService)
	})
}