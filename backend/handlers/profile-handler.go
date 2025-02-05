package handlers

import (
	"kfs-backend/services"
	"net/http"

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

// Profil oluşturma isteği için yapı
type CreateProfileRequest struct {
	PhotoURL       string `json:"photoUrl"`       // Fotoğraf URL
	Website        string `json:"website"`        // Kişisel web sitesi
	IdentityNumber string `json:"identityNumber"` // TC Kimlik No
	BirthDate      string `json:"birthDate"`      // Doğum tarihi
	Gender         string `json:"gender"`         // Cinsiyet
	AcademicTitle  string `json:"academicTitle"`  // Akademik unvan
}

// Profil bilgilerini getirir
func GetProfileByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId") // URL'den userId'yi alır

	profile, err := services.GetProfileByUserId(userId)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Profil başarıyla güncellendi",
		"profile": profile,
	})
}

// Yeni profil oluşturur
func CreateProfile(c *fiber.Ctx) error {
	var req CreateProfileRequest

	// İstek gövdesini parse et
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz istek formatı",
		})
	}

	userId := c.Locals("userId").(uint)

	profile, err := services.CreateProfile(
		userId,
		req.PhotoURL,
		req.Website,
		req.IdentityNumber,
		req.BirthDate,
		req.Gender,
		req.AcademicTitle,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Profil başarıyla oluşturuldu",
		"profile": profile,
	})
}
