package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupVentureSectorRoutes(app *fiber.App) {
	// Generic service'i VentureSector için başlat
	sectorService := &services.GenericVentureService[models.VentureSector]{}

	// Route gruplarını oluştur
	sectorGroup := app.Group("/api/venture-sectors")

	// Yeni VentureSector oluştur
	sectorGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.VentureSectorRequest
		return handlers.CreateVenture(c, sectorService, req)
	})

	// ID ile VentureSector getir
	sectorGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, sectorService)
	})

	// Campaign ID'ye göre VentureSector listele
	sectorGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, sectorService, "campaign_id")
	})

	// VentureSector güncelle
	sectorGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.VentureSectorRequest
		return handlers.UpdateVenture(c, sectorService, req)
	})

	// VentureSector sil
	sectorGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, sectorService)
	})
}
