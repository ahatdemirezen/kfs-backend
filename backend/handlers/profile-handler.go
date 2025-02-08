package handlers

import (
	"kfs-backend/services"

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
	var req UpdateProfileRequest

	// İstek gövdesini parse et
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Geçersiz istek formatı")
	}

	userId := c.Locals("userId").(uint)

	profile, err := services.UpdateProfile(
		userId,
		req.PhotoURL,
		req.Website,
		req.IdentityNumber,
		req.BirthDate,
		req.Gender,
		req.AcademicTitle,
	)

	if err != nil {
		return err // Service'den gelen fiber.Error'u direkt olarak dön
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profil başarıyla güncellendi",
		"profile": profile,
	})
}
