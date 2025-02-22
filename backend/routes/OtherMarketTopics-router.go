package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupOtherMArketTopicsRoutes(app *fiber.App) {
	otherMarketTopicsService := &services.GenericVentureService[models.OtherMarketTopic]{} // HL

	// Route gruplarını olustur
	otherMarketTopicsGroup := app.Group("/api/other-market-topics")

	// Yeni OtherMarketTopics olustur
	otherMarketTopicsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.OtherMarketTopicRequest
		return handlers.CreateVenture(c, otherMarketTopicsService, req)
	})

	// ID ile OtherMarketTopics getir
	otherMarketTopicsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, otherMarketTopicsService)
	})

	// Campaign ID'ye gore OtherMarketTopics listele
	otherMarketTopicsGroup.Get("/list/:market_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, otherMarketTopicsService, "market_info_id")
	})

	// OtherMarketTopics guncelle
	otherMarketTopicsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.OtherMarketTopicRequest
		return handlers.UpdateVenture(c, otherMarketTopicsService, req)
	})

	// OtherMarketTopics sil
	otherMarketTopicsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, otherMarketTopicsService)
	})
}