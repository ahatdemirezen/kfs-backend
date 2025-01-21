package main

import (
	"log"

	"kfs-backend/database" // Veritabanı bağlantısı

	"github.com/gofiber/fiber/v2" // Fiber framework
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Fiber uygulamasını başlat
	app := fiber.New()

	// CORS middleware'ini ekle
	app.Use(cors.New())

	// Veritabanı bağlantısını başlat
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fonbulucu API'sine Hoş Geldiniz")
	})
	// Rotaları tanımla

	// Uygulamayı başlat
	log.Fatal(app.Listen(":3000"))
}
