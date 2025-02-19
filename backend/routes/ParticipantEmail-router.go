package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupParticipantEmailRoutes(app *fiber.App) {
	// Generic service'i ParticipantEmail için başlat
	participantEmailService := &services.GenericVentureService[models.ParticipantEmail]{}

	// Route gruplarını oluştur
	participantEmailGroup := app.Group("/api/participant-emails")

	// Yeni ParticipantEmail oluştur
	participantEmailGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ParticipantEmailRequest
		return handlers.CreateVenture(c, participantEmailService, req)
	})

	// ID ile ParticipantEmail getir
	participantEmailGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, participantEmailService)
	})

	// Campaign ID'ye göre ParticipantEmail listele
	participantEmailGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, participantEmailService, "campaign_id")
	})

	// ParticipantEmail güncelle
	participantEmailGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ParticipantEmailRequest
		return handlers.UpdateVenture(c, participantEmailService, req)
	})

	// ParticipantEmail sil
	participantEmailGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, participantEmailService)
	})
}
