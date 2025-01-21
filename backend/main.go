package main

import (
	"log"

	"kfs-backend/repositories" // Veritabanı bağlantısı

	"github.com/gofiber/fiber/v2" // Fiber framework
)

func main() {
	// Fiber uygulamasını başlat
	app := fiber.New()

	// Veritabanı bağlantısını başlat
	repositories.ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	// Rotaları tanımla

	// Uygulamayı başlat
	log.Fatal(app.Listen(":3000"))
}
