package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupPatentRoutes(app *fiber.App) {
	// Generic service'i Patent için baslat
	patentService := &services.GenericVentureService[models.Patent]{} // HL

	// Route gruplarını olustur
	patentGroup := app.Group("/api/patents")

	// Yeni Patent olustur
	patentGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.PatentRequest
		return handlers.CreateVenture(c, patentService, req)
	})

	// ID ile Patent getir
	patentGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, patentService)
	})	

	// Campaign ID'ye gore Patent listele
	patentGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, patentService, "campaign_id")
	})

	// Patent guncelle
	patentGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.PatentRequest
		return handlers.UpdateVenture(c, patentService, req)
	})

	// Patent sil
	patentGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, patentService)
	})
}