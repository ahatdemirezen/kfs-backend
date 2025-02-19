package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupExtraFinancingResourcesRoutes(app *fiber.App) {
	// Generic service'i ExtraFinancingResource için başlat
	extraFinancingResourceService := &services.GenericVentureService[models.ExtraFinancingResource]{}

	// Route gruplarını oluştur
	extraFinancingResourcesGroup := app.Group("/api/extra-financing-resources")

	// Yeni ExtraFinancingResource oluştur
	extraFinancingResourcesGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.ExtraFinancingResourceRequest
		return handlers.CreateVenture(c, extraFinancingResourceService, req)
	})

	// ID ile ExtraFinancingResource getir
	extraFinancingResourcesGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, extraFinancingResourceService)
	})

	// FundingInfoId'ye göre ExtraFinancingResource listele
	extraFinancingResourcesGroup.Get("/list/:funding_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, extraFinancingResourceService , "funding_info_id")
	})

	// ExtraFinancingResource güncelle
	extraFinancingResourcesGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.ExtraFinancingResourceRequest
		return handlers.UpdateVenture(c, extraFinancingResourceService, req)
	})

	// ExtraFinancingResource sil
	extraFinancingResourcesGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, extraFinancingResourceService)
	})
}
