package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupOtherProductTopicsRoutes(app *fiber.App) {
	// Generic service'i OtherProductTopics için baslat
	otherProductTopicsService := &services.GenericVentureService[models.OtherProductTopic]{} // HL

	// Route gruplarını olustur
	otherProductTopicsGroup := app.Group("/api/other-product-topics")

	// Yeni OtherProductTopics olustur
	otherProductTopicsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.OtherProductTopicRequest
		return handlers.CreateVenture(c, otherProductTopicsService, req)
	})

	// ID ile OtherProductTopics getir
	otherProductTopicsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, otherProductTopicsService)
	})

	// Campaign ID'ye gore OtherProductTopics listele
	otherProductTopicsGroup.Get("/list/:product_model_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, otherProductTopicsService, "product_model_info_id")
	})

	// OtherProductTopics guncelle
	otherProductTopicsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.OtherProductTopicRequest
		return handlers.UpdateVenture(c, otherProductTopicsService, req)
	})

	// OtherProductTopics sil
	otherProductTopicsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, otherProductTopicsService)
	})
}