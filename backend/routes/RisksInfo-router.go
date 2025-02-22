package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRisksInfoRoutes(app *fiber.App) {
	// Generic service'i RisksInfo için başlat
	risksInfoService := &services.GenericVentureService[models.RisksInfo]{} // HL

	// Route gruplarını oluştur
	risksInfoGroup := app.Group("/api/risks-info")

	// Yeni RisksInfo oluştur
	risksInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.RisksInfoRequest
		return handlers.CreateVenture(c, risksInfoService, req)
	})

	// ID ile RisksInfo getir
	risksInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, risksInfoService)
	})

	// CampaignId'ye göre RisksInfo listele
	risksInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, risksInfoService, "campaign_id")
	})

	// RisksInfo güncelle
	risksInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.RisksInfoRequest
		return handlers.UpdateVenture(c, risksInfoService, req)
	})

	// RisksInfo sil
	risksInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, risksInfoService)
	})
}
