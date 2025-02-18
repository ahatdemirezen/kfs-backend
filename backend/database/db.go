package database

import (
	"fmt"
	"log"
	"strings"

	"kfs-backend/config" // Config paketini içe aktar

	"kfs-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Tabloları migrate et
func migrateIfNotExists(db *gorm.DB) error {
	// Migrasyon sırasını belirle
	migrations := []struct {
		Model interface{}
		Name  string
	}{
		{&models.User{}, "users"},
		{&models.Profile{}, "profiles"},
		{&models.Verification{}, "verifications"},
		{&models.Role{}, "roles"},
	}

	// Her model için ayrı ayrı migrasyon yap
	for _, migration := range migrations {
		err := db.AutoMigrate(migration.Model)
		if err != nil {
			// Eğer tablo zaten varsa veya benzer bir hata ise devam et
			if strings.Contains(err.Error(), "already exists") {
				log.Printf("%s tablosu zaten mevcut, geçiliyor...", migration.Name)
				continue
			}
			// Diğer hataları raporla
			log.Printf("%s tablosu için migrasyon hatası: %v", migration.Name, err)
		} else {
			log.Printf("%s tablosu başarıyla migrate edildi", migration.Name)
		}
	}

	session := db.Session(&gorm.Session{PrepareStmt: true})
	if session.Error != nil {
		fmt.Println("Session error: ", session.Error)
	} else {
		fmt.Println("Session created successfully")
	}
	return nil
}

// tryconnectingdb
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

	// GORM konfigürasyonu
	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt: true, // Prepared statement'ları devre dışı bırak
		Logger: logger.Default.LogMode(logger.Silent), // SQL loglarını kapat
		SkipDefaultTransaction: true, // Varsayılan transaction'ları devre dışı bırak
	}

	// Veritabanına bağlan
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı: ", err)
	}
	log.Println("Veritabanına bağlantı başarılı")

	// Migrasyon işlemini gerçekleştir
	if err := migrateIfNotExists(db); err != nil {
		log.Printf("Migrasyon sırasında beklenmeyen bir hata oluştu: %v", err)
	} else {
		log.Println("Tüm migrasyon işlemleri tamamlandı")
	}

	DB = db
}
