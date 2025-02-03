package handlers

import (
	"kfs-backend/database"
	"kfs-backend/models"
	"net/http"
	"time" // Zaman işlemleri için gerekli paket

	"github.com/gofiber/fiber/v2"
)

// Kullanıcının profil bilgilerini güncelleme talebi için bir yapı
type UpdateProfileRequest struct {
	UserId         uint   `json:"userId" validate:"required"` // User ID
	PhotoURL       string `json:"photoUrl"`                   // Fotoğraf URL
	Website        string `json:"website"`                    // Kişisel web sitesi
	IdentityNumber string `json:"identityNumber"`             // TC Kimlik No
	BirthDate      string `json:"birthDate"`                  // Doğum tarihi
	Gender         string `json:"gender"`                     // Cinsiyet
	AcademicTitle  string `json:"academicTitle"`              // Akademik unvan
}

// Profil bilgilerini getirir
func GetProfileByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId") // URL'den userId'yi alır

	var profile models.Profile
	// Veritabanında user_id'ye göre profil arar
	if err := database.DB.Where("user_id = ?", userId).First(&profile).Error; err != nil {
		// Profil bulunamazsa 404 döner
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Profil bulunamadı",
		})
	}

	// Profil bilgilerini JSON formatında döner
	return c.Status(http.StatusOK).JSON(profile)
}

func UpdateProfile(c *fiber.Ctx) error {
	var req UpdateProfileRequest

	// İstek gövdesini parse et
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz istek formatı",
		})
	}

	var profile models.Profile
	// Kullanıcının profilini bul
	if err := database.DB.Where("user_id = ?", req.UserId).First(&profile).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Profil bulunamadı",
		})
	}

	// Doğum tarihini string'den time.Time formatına dönüştür
	birthDate, err := time.Parse("2006-01-02", req.BirthDate) // YYYY-MM-DD formatını kullanıyoruz
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Doğum tarihi formatı geçersiz. Beklenen format: YYYY-MM-DD",
		})
	}

	// Profili güncelle
	profile.PhotoURL = req.PhotoURL
	profile.Website = req.Website
	profile.IdentityNumber = req.IdentityNumber
	profile.BirthDate = birthDate // Dönüştürülmüş `time.Time` değeri
	profile.Gender = req.Gender
	profile.AcademicTitle = req.AcademicTitle

	// Güncellemeyi veritabanına kaydet
	if err := database.DB.Save(&profile).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Profil güncellenirken bir hata oluştu",
		})
	}

	// Güncellenmiş profili döner
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Profil başarıyla güncellendi",
		"profile": profile,
	})
}
