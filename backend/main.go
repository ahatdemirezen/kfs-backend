package main

import (
	"log"

	"kfs-backend/config"
	"kfs-backend/database" // Veritabanı bağlantısı

	"github.com/gofiber/fiber/v2" // Fiber framework
	"github.com/gofiber/fiber/v2/middleware/cors"
	"kfs-backend/routes"
)

func main() {
	// Fiber uygulamasını başlat
	app := fiber.New()

	// CORS middleware'ini ekle
	app.Use(cors.New())

	// .env yükle (JWT secret vb.)
	config.LoadConfig()

	// Veritabanı bağlantısını baş	lat
	database.ConnectDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Fonbulucu API'sine Hoş Geldiniz")
	})
	// Rotaları tanımla
	routes.SetupUserRoutes(app)
	routes.SetupProfileRoutes(app) // Profil rotaları
	routes.SetupAuthRoutes(app) //auth route'larını kaydet
	// Uygulamayı başlat
	log.Fatal(app.Listen(":3000"))
}