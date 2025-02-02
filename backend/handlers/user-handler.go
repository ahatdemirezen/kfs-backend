	package handlers

	import (
		"github.com/gofiber/fiber/v2"
		"golang.org/x/crypto/bcrypt"
		"kfs-backend/database"
		"kfs-backend/models"
		"crypto/rand"
		"encoding/hex"
		"strconv"
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

		// Doğrulama kodunu kontrol et
		if !VerifyCode(req.Email, req.Code) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Geçersiz doğrulama kodu",
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
			Email:    req.Email,
			Password: string(hashedPassword),
			Salt:     salt,
		}

		// Kullanıcıyı veritabanına kaydet
		if result := db.Create(&user); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Kullanıcı oluşturulurken bir hata oluştu",
				"details": result.Error.Error(),
			})
		}

		// Verification kaydı oluştur
		verification := models.Verification{
			UserId:          user.UserId,
			IsEmailVerified: true, // Email doğrulandığı için true
			IsPhoneVerified: false,
			IsUserVerified:  false,
			IsLawApproved:   false,
		}

		if result := db.Create(&verification); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Doğrulama kaydı oluşturulurken bir hata oluştu",
				"details": result.Error.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Kullanıcı başarıyla oluşturuldu ve email doğrulandı",
			"userId": user.UserId,
			"user": user,
		})
	}
	func UpdateUser(c *fiber.Ctx) error {
		db := database.DB

		// URL'den gelen userId parametresini al
		userIDParam := c.Params("userId")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Geçersiz kullanıcı ID",
			})
		}

		// Kullanıcıyı veritabanında ara
		var user models.User
		if result := db.First(&user, userID); result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Kullanıcı bulunamadı",
			})
		}

		// İstekten gelen güncelleme verilerini parse et
		var req UpdateUserRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Geçersiz istek formatı",
			})
		}

		// Kullanıcı bilgilerini güncelle
		user.FirstName = req.FirstName
		user.LastName = req.LastName
		user.Phone = req.Phone
		user.Country = req.Country
		user.CompanyName = req.CompanyName
		user.TaxOffice = req.TaxOffice
		user.TaxNumber = req.TaxNumber

		// Veritabanında güncelle
		if result := db.Save(&user); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Kullanıcı güncellenirken hata oluştu",
				"details": result.Error.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Kullanıcı başarıyla güncellendi",
			"user": user,
		})
	}
