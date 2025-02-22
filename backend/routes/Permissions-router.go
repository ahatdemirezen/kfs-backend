package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupPermissionsRoutes(app *fiber.App) {
	// Generic service'i Permissions için baslat
	permissionsService := &services.GenericVentureService[models.Permission]{} // HL

	// Route gruplarını olustur
	permissionsGroup := app.Group("/api/permissions")	

	// Yeni Permissions olustur
	permissionsGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.PermissionRequest
		return handlers.CreateVenture(c, permissionsService, req)
	})

	// ID ile Permissions getir
	permissionsGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, permissionsService)
	})

	// Campaign ID'ye gore Permissions listele
	permissionsGroup.Get("/list/:campaign_id", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, permissionsService, "campaign_id")
	})
	
	// Permissions guncelle
	permissionsGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.PermissionRequest
		return handlers.UpdateVenture(c, permissionsService, req)
	})

	// Permissions sil
	permissionsGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, permissionsService)
	})
}