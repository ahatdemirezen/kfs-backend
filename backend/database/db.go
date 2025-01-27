package database

import (
	"fmt"
	"log"

	"kfs-backend/config" // Config paketini içe aktar
	"kfs-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Config paketinden ayarları al
	cfg := config.AppConfig

	// DSN oluştur
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		cfg.SupabaseHost,
		cfg.SupabaseUser,
		cfg.SupabasePassword,
		cfg.SupabaseDBName,
		cfg.SupabasePort,
	)

	// Veritabanına bağlan
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı: ", err)
	}
	log.Println("Veritabanına bağlantı başarılı")

	// Tabloları otomatik oluştur
	err = db.AutoMigrate(&models.User{}, &models.Profile{})
	if err != nil {
		log.Fatal("Veritabanı tabloları oluşturulamadı: ", err)
	}

	DB = db
}
