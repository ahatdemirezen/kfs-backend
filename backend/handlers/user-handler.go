package handlers

import (
	"kfs-backend/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Code     string `json:"code" validate:"required"` // Doğrulama kodu
}

type VerifyEmailRequest struct {
	UserId uint   `json:"userId" validate:"required"`
	Token  string `json:"token" validate:"required"`
}

type UpdateUserRequest struct {
	UserId        uint   `json:"userId" validate:"required"`
	FirstName     string `json:"firstName" validate:"required"`
	LastName      string `json:"lastName" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	Country       string `json:"country" validate:"required"`
	CompanyName   string `json:"companyName" validate:"required"`
	TaxOffice     string `json:"taxOffice" validate:"required"`
	TaxNumber     string `json:"taxNumber" validate:"required"`
	IsLawApproved bool   `json:"isLawApproved" validate:"required"`
}

func Register(c *fiber.Ctx) error {
	// Request body'i parse et
	var req services.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "İstek formatı geçersiz")
	}

	// Doğrulama kodunu kontrol et
	if !VerifyCode(req.Email, req.Code) {
		return fiber.NewError(fiber.StatusBadRequest, "Geçersiz doğrulama kodu")
	}

	user, err := services.RegisterUser(req)
	if err != nil {
		return err // Service'den gelen fiber.Error'u direkt olarak dön
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kullanıcı başarıyla oluşturuldu ve email doğrulandı",
		"userId":  user.UserId,
		"user":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	// URL'den gelen userId parametresini al
	userIDParam := c.Params("userId")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Geçersiz kullanıcı ID")
	}

	// Query'den userType parametresini al
	userType := c.Query("userType")
	if userType == "" || (userType != "individual" && userType != "corporate") {
		return fiber.NewError(fiber.StatusBadRequest, "Geçersiz userType değeri. 'individual' veya 'corporate' olmalıdır")
	}

	// İstekten gelen güncelleme verilerini parse et
	var req services.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Geçersiz istek formatı")
	}

	user, verification, err := services.UpdateUserInfo(uint(userID), userType, req)
	if err != nil {
		return err // Service'den gelen fiber.Error'u direkt olarak dön
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Kullanıcı ve rol başarıyla güncellendi",
		"user":         user,
		"verification": verification,
	})
}

func GetUser(c *fiber.Ctx) error {
	// Middleware'den userId'yi al
	userId := c.Locals("userId").(uint)

	profile, err := services.GetUser(userId)
	if err != nil {
		return err // Service'den gelen fiber.Error'u direkt olarak dön
	}

	return c.Status(fiber.StatusOK).JSON(profile)
}
