package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupVentureBusinessModalRoutes(app *fiber.App) {
	// Generic service'i VentureBusinessModal için başlat
	businessModalService := &services.GenericVentureService[models.VentureBusinessModal]{}

	// Route gruplarını oluştur
	businessModalGroup := app.Group("/api/venture-business-modals")

	// Yeni VentureBusinessModal oluştur
	businessModalGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.VentureBusinessModalRequest
		return handlers.CreateVenture(c, businessModalService, req)
	})

	// ID ile VentureBusinessModal getir
	businessModalGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, businessModalService)
	})

	// Campaign ID'ye göre VentureBusinessModal listele
	businessModalGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, businessModalService, "campaign_id")
	})

	// VentureBusinessModal güncelle
	businessModalGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.VentureBusinessModalRequest
		return handlers.UpdateVenture(c, businessModalService, req)
	})

	// VentureBusinessModal sil
	businessModalGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, businessModalService)
	})
}
