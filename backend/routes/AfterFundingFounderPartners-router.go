package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupAfterFundingFounderPartnersRoutes(app *fiber.App) {
	// Initialize generic service for AfterFundingFounderPartners
	afterFundingFounderPartnersService := &services.GenericVentureService[models.AfterFundingFounderPartner]{}

	// Create route group
	afterFundingFounderPartnersGroup := app.Group("/api/after-funding-founder-partners")

	// Create new AfterFundingFounderPartner
	afterFundingFounderPartnersGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.AfterFundingFounderPartnerRequest
		return handlers.CreateVenture(c, afterFundingFounderPartnersService, req)
	})

	// Retrieve AfterFundingFounderPartner by ID
	afterFundingFounderPartnersGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, afterFundingFounderPartnersService)
	})

	// List AfterFundingFounderPartners by EnterpriseInfoId
	afterFundingFounderPartnersGroup.Get("/list/:enterprise_info_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, afterFundingFounderPartnersService, "enterprise_info_id")
	})

	// Update AfterFundingFounderPartner
	afterFundingFounderPartnersGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.AfterFundingFounderPartnerRequest
		return handlers.UpdateVenture(c, afterFundingFounderPartnersService, req)
	})

	// Delete AfterFundingFounderPartner
	afterFundingFounderPartnersGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, afterFundingFounderPartnersService)
	})
}
