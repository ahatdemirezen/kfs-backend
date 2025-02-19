package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupAchivementsRoutes(app *fiber.App) {
	// Generic service'i Achivements için baslat
	achivementsService := &services.GenericVentureService[models.Achievement]{} // HL

	// Route gruplarını olustur
	achivementsGroup := app.Group("/api/achivements")

	// Yeni Achivements olustur
	achivementsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.AchievementRequest
		return handlers.CreateVenture(c, achivementsService, req)
	})

	// ID ile Achivements getir
	achivementsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, achivementsService)
	})	

	achivementsGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, achivementsService, "campaign_id")
	})

	// Achivements guncelle
	achivementsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.AchievementRequest
		return handlers.UpdateVenture(c, achivementsService, req)
	})

	// Achivements sil
	achivementsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, achivementsService)
	})
}