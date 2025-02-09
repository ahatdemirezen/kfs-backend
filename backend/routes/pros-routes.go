package routes

import (
	"kfs-backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupProsRoutes(app *fiber.App) {
	pros := app.Group("/api/pros")

	pros.Post("/", handlers.CreatePro)
	pros.Get("/:proId", handlers.GetPro)
	pros.Get("/analysis/:analysisInfoId", handlers.GetProsByAnalysisInfo)
	pros.Put("/:proId", handlers.UpdatePro)
	pros.Delete("/:proId", handlers.DeletePro)
	pros.Post("/bulk", handlers.CreateMultiplePros)

}
