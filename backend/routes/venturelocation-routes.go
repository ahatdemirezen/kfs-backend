package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupVentureLocationRoutes(app *fiber.App) {
	// Generic service'i VentureLocation için başlat
	locationService := &services.GenericVentureService[models.VentureLocation]{}

	// Route gruplarını oluştur
	locationGroup := app.Group("api/venture-locations")

	// Yeni VentureLocation oluştur
	locationGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.VentureLocationRequest
		return handlers.CreateVenture(c, locationService, req)
	})

	// ID ile VentureLocation getir
	locationGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, locationService)
	})

	// Campaign ID'ye göre VentureLocation listele
	locationGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, locationService, "campaign_id")
	})

	// VentureLocation güncelle
	locationGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.VentureLocationRequest
		return handlers.UpdateVenture(c, locationService, req)
	})

	// VentureLocation sil
	locationGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, locationService)
	})
}