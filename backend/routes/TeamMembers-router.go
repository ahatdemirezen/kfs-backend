package routes

import (
	"kfs-backend/handlers"
	"kfs-backend/models"
	"kfs-backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupTeamMembersRoutes(app *fiber.App) {
	// Generic service'i TeamMembers için baslat
	teamMembersService := &services.GenericVentureService[models.TeamMember]{} // HL

	// Route gruplarını olustur
	teamMembersGroup := app.Group("/api/team-members")

	// Yeni TeamMembers olustur
	teamMembersGroup.Post("/create", func(c *fiber.Ctx) error {
		var req services.TeamMemberRequest
		return handlers.CreateVenture(c, teamMembersService, req)
	})

	// ID ile TeamMembers getir
	teamMembersGroup.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetVentureByID(c, teamMembersService)
	})	

	// Campaign ID'ye gore TeamMembers listele
	teamMembersGroup.Get("/list/:campaignId", func(c *fiber.Ctx) error {
		return handlers.GetVenturesByField(c, teamMembersService, "campaign_id")
	})

	// TeamMembers guncelle
	teamMembersGroup.Put("/update/:id", func(c *fiber.Ctx) error {
		var req services.TeamMemberRequest
		return handlers.UpdateVenture(c, teamMembersService, req)
	})

	// TeamMembers sil
	teamMembersGroup.Delete("/delete/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteVenture(c, teamMembersService)
	})
}