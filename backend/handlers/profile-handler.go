package handlers

import (
	"kfs-backend/services"
	"mime/multipart"

	// Zaman işlemleri için gerekli paket
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

func UpdateProfile(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	// Form-data içinden dosya al
	var file *multipart.FileHeader
	file, err := c.FormFile("photo")
	if err != nil {
		file = nil // Dosya yüklenmemişse, nil olarak ayarla
	}

	// Diğer form alanlarını al
	website := c.FormValue("website")
	identityNumber := c.FormValue("identityNumber")
	birthDate := c.FormValue("birthDate")
	gender := c.FormValue("gender")
	academicTitle := c.FormValue("academicTitle")

	// Profili güncelle
	profile, err := services.UpdateProfile(userId, file, website, identityNumber, birthDate, gender, academicTitle)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profil başarıyla güncellendi",
		"profile": profile,
	})
}

func UpdateProfilePhoto(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	// Form-data içinden dosya al
	file, err := c.FormFile("photo")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Dosya yüklenemedi, lütfen bir dosya seçin")
	}

	// **Sadece profil fotoğrafını güncelle**
	profile, err := services.UpdateProfilePhoto(userId, file)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profil fotoğrafı başarıyla güncellendi",
		"profile": profile,
	})
}
