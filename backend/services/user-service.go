package services

import (
	"crypto/rand"
	"encoding/hex"
	"kfs-backend/database"
	"kfs-backend/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Code     string `json:"code" validate:"required"`
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

func generateSalt() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}

func RegisterUser(req RegisterRequest) (*models.User, error) {
	db := database.DB

	// Email kontrolü
	var existingUser models.User
	if result := db.Where("email = ?", req.Email).First(&existingUser); result.Error == nil {
		return nil, result.Error
	}

	// Salt oluştur
	salt := generateSalt()

	// Şifreyi hashle
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password+salt), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Yeni kullanıcı oluştur
	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Salt:     salt,
	}

	// Kullanıcıyı veritabanına kaydet
	if result := db.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	// Verification kaydı oluştur
	verification := models.Verification{
		UserId:          user.UserId,
		IsEmailVerified: true,
		IsPhoneVerified: false,
		IsUserVerified:  false,
		IsLawApproved:   false,
	}

	if result := db.Create(&verification); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func UpdateUserInfo(userID uint, userType string, req UpdateUserRequest) (*models.User, *models.Verification, error) {
	db := database.DB

	// Kullanıcıyı veritabanında ara
	var user models.User
	if result := db.First(&user, userID); result.Error != nil {
		return nil, nil, result.Error
	}

	var verification models.Verification
	if result := db.Where("user_id = ?", userID).First(&verification); result.Error != nil {
		return nil, nil, result.Error
	}

	// Kullanıcı bilgilerini güncelle
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Phone = req.Phone
	user.Country = req.Country
	user.CompanyName = req.CompanyName
	user.TaxOffice = req.TaxOffice
	user.TaxNumber = req.TaxNumber
	verification.IsLawApproved = req.IsLawApproved

	// Veritabanında güncelle
	if result := db.Save(&user); result.Error != nil {
		return nil, nil, result.Error
	}

	if result := db.Save(&verification); result.Error != nil {
		return nil, nil, result.Error
	}

	// Profil kontrolü yap, yoksa oluştur
	var profile models.Profile
	result := db.Where("user_id = ?", userID).First(&profile)
	if result.Error != nil {
		// Profil bulunamadıysa yeni profil oluştur
		newProfile := models.Profile{
			UserId: userID,
		}
		if err := db.Create(&newProfile).Error; err != nil {
			return nil, nil, err
		}
	}

	// Kullanıcının mevcut rolünü kontrol et
	var existingRole models.Role
	result = db.Where("user_id = ?", userID).First(&existingRole)

	if result.Error != nil {
		// Rol bulunamadıysa yeni rol oluştur
		newRole := models.Role{
			UserId: userID,
			Role:   userType,
		}
		if err := db.Create(&newRole).Error; err != nil {
			return nil, nil, err
		}
	} else {
		// Mevcut rolü güncelle
		existingRole.Role = userType
		if err := db.Save(&existingRole).Error; err != nil {
			return nil, nil, err
		}
	}

	return &user, &verification, nil
}

// Kullanıcı ve profil bilgilerini getir
func GetUser(userId uint) (*models.Profile, error) {
	db := database.DB
	var profile models.Profile

	// Kullanıcının profilini ve ilişkili user bilgilerini getir
	if result := db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("user_id, email, country, phone, first_name, last_name, company_name, tax_office, tax_number, created_at, updated_at")
	}).Where("user_id = ?", userId).First(&profile); result.Error != nil {
		// Profil bulunamazsa yeni profil oluştur
		newProfile := models.Profile{
			UserId: userId,
		}
		if err := db.Create(&newProfile).Error; err != nil {
			return nil, err
		}
		// Yeni oluşturulan profilin user bilgilerini de getir
		if err := db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, email, country, phone, first_name, last_name, company_name, tax_office, tax_number, created_at, updated_at")
		}).First(&profile, newProfile.ProfileId).Error; err != nil {
			return nil, err
		}
	}

	return &profile, nil
}
