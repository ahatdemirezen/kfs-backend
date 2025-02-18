package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"kfs-backend/database"
	"kfs-backend/models"
	"kfs-backend/services"
	"crypto/rand"
	"encoding/hex"
	"time"
)

type RegisterRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6"`
	Country     string `json:"country"`
	Phone       string `json:"phone"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	CompanyName string `json:"companyName"`
	TaxOffice   string `json:"taxOffice"`
	TaxNumber   string `json:"taxNumber"`
}


type VerifyEmailRequest struct {
	UserId uint   `json:"userId" validate:"required"`
	Token  string `json:"token" validate:"required"`
}

type UpdateUserRequest struct {
	UserId      uint   `json:"userId" validate:"required"`
	FirstName   string `json:"firstName" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Country     string `json:"country" validate:"required"`
	CompanyName string `json:"companyName" validate:"required"`
	TaxOffice   string `json:"taxOffice" validate:"required"`
	TaxNumber   string `json:"taxNumber" validate:"required"`
}

func generateSalt() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func Register(c *fiber.Ctx) error {
	db := database.DB
	
	// Request body'i parse et
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "İstek formatı geçersiz",
		})
	}

	// Email kontrolü
	var existingUser models.User
	if result := db.Where("email = ?", req.Email).First(&existingUser); result.Error == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Bu email adresi zaten kayıtlı",
		})
	}

	// Salt oluştur
	salt := generateSalt()

	// Şifreyi hashle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Şifre işlenirken bir hata oluştu",
		})
	}

	// Yeni kullanıcı oluştur
	user := models.User{
		Email:       req.Email,
		Password:    string(hashedPassword),
		Salt:        salt,
		Country:     req.Country,
		Phone:       req.Phone,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		CompanyName: req.CompanyName,
		TaxOffice:   req.TaxOffice,
		TaxNumber:   req.TaxNumber,
	}

	// Kullanıcıyı veritabanına kaydet
	if result := db.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Kullanıcı oluşturulurken bir hata oluştu",
			"details": result.Error.Error(),
		})
	}

	// Doğrulama kodu oluştur
	verificationCode := services.GenerateVerificationCode()
	
	// Verification kaydı oluştur
	verification := models.Verification{
		UserId:          user.UserId,
		IsEmailVerified: false,
		IsPhoneVerified: false,
		IsUserVerified:  false,
		IsLawApproved:   false,
		EmailVerificationCode: verificationCode,
		EmailCodeExpiry: time.Now().Add(15 * time.Minute).Unix(), // 15 dakika geçerli
	}

	if result := db.Create(&verification); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Doğrulama kaydı oluşturulurken bir hata oluştu",
			"details": result.Error.Error(),
		})
	}

	// Email gönder
	if err := services.SendVerificationEmail(user.Email, verificationCode); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Doğrulama e-postası gönderilirken bir hata oluştu",
			"details": err.Error(),
		})
	}

	// Gelen userType'e göre role belirle
	userType := c.Query("userType")
	var roleValue string
	if userType == "bireysel" {
		roleValue = "bireysel"
	} else if userType == "kurumsal" {
		roleValue = "kurumsal"
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz userType. Bireysel veya kurumsal olmalıdır.",
		})
	}

	// Role kaydı oluştur
	role := models.Role{
		UserId: user.UserId,
		Role:   roleValue,
	}

	if result := db.Create(&role); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Role kaydı oluşturulurken bir hata oluştu",
			"details": result.Error.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kullanıcı başarıyla oluşturuldu. Lütfen email adresinize gönderilen kodu kullanarak hesabınızı doğrulayın.",
		"userId": user.UserId,
		"user": user,
		"role": role,
	})
}

func VerifyEmail(c *fiber.Ctx) error {
	db := database.DB

	var req VerifyEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "İstek formatı geçersiz",
		})
	}

	var verification models.Verification
	if result := db.Where("user_id = ?", req.UserId).First(&verification); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Kullanıcı bulunamadı",
		})
	}

	// Kodun süresinin dolup dolmadığını kontrol et
	if time.Now().Unix() > verification.EmailCodeExpiry {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Doğrulama kodunun süresi dolmuş",
		})
	}

	// Doğrulama kodunu kontrol et
	if verification.EmailVerificationCode != req.Token {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Geçersiz doğrulama kodu",
		})
	}

	// Email doğrulamasını güncelle
	verification.IsEmailVerified = true
	if result := db.Save(&verification); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Doğrulama güncellenirken bir hata oluştu",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Email adresi başarıyla doğrulandı",
	})
}
