package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupUsageRoutes(app *fiber.App) {
	// Generic service'i Usage için başlat
	usageService := &services.GenericVentureService[models.Usage]{} // HL

	// Route gruplarını oluştur
	usageGroup := app.Group("/api/usages")

	// Yeni Usage oluştur
	usageGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.UsageRequest
		return handlers.CreateVenture(c, usageService, req)
	})

	// ID ile Usage getir
	usageGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, usageService)
	})

	// FundingInfoId'ye göre Usage listele
	usageGroup.Get("/list/:funding_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, usageService, "funding_info_id")
	})

	// Usage güncelle
	usageGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.UsageRequest
		return handlers.UpdateVenture(c, usageService, req)
	})

	// Usage sil
	usageGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, usageService)
	})
}
