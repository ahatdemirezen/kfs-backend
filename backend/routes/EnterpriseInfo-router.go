package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupEnterpriseInfoRoutes(app *fiber.App) {
	// Generic service'i EnterpriseInfo için başlat
	enterpriseInfoService := &services.GenericVentureService[models.EnterpriseInfo]{}

	// Route gruplarını oluştur
	enterpriseInfoGroup := app.Group("/api/enterprise-info")

	// Yeni EnterpriseInfo oluştur
	enterpriseInfoGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.EnterpriseInfoRequest
		return handlers.CreateVenture(c, enterpriseInfoService, req)
	})

	// ID ile EnterpriseInfo getir
	enterpriseInfoGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, enterpriseInfoService)
	})

	// CampaignId'ye göre EnterpriseInfo listele
	enterpriseInfoGroup.Get("/list/:campaign_Id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, enterpriseInfoService, "campaign_id")
	})

	// EnterpriseInfo güncelle
	enterpriseInfoGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.EnterpriseInfoRequest
		return handlers.UpdateVenture(c, enterpriseInfoService, req)
	})

	// EnterpriseInfo sil
	enterpriseInfoGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, enterpriseInfoService)
	})
}
