package utils

import (
	"log"
	"kfs-backend/database"
	"kfs-backend/models"
	"strconv"
	"github.com/gofiber/fiber/v2"
	"errors"
)


// **ID'yi Al ve uint'e Çevir**
// func GetIDParam(c *fiber.Ctx) (uint, error) {
// 	id, err := strconv.Atoi(c.Params("id")) // <- Burada gelen parametre "" olabilir
// 	if err != nil {
// 		return 0, err
// 	}
// 	return uint(id), nil
// }


// func GetIDParam(c *fiber.Ctx) (uint, error) {
// 	// **Öncelikle "campaignId" olup olmadığını kontrol et**
// 	idParam := c.Params("campaignId") // Eğer "/list/:campaignId" varsa buradan alır
// 	if idParam == "" {
// 		// **Eğer "campaignId" yoksa, "id" parametresini al**
// 		idParam = c.Params("id") // Eğer "/get/:id" gibi bir route varsa buradan alır
// 	}

// 	// **Log ekleyerek hangi parametrenin alındığını kontrol et**
// 	log.Println("Gelen ID Parametresi:", idParam)

// 	// **Eğer hala boşsa hata döndür**
// 	if idParam == "" {
// 		log.Println("Hata: ID parametresi eksik!")
// 		return 0, errors.New("ID parametresi eksik")
// 	}

// 	// **ID'yi uint'e çevir**
// 	id, err := strconv.ParseUint(idParam, 10, 32)
// 	if err != nil {
// 		log.Println("Hata: ID uint'e çevrilemedi -", err.Error())
// 		return 0, err
// 	}

// 	return uint(id), nil
// }


// GetIDParam, dinamik olarak parametre adını alır ve uint'e çevirir.
func GetIDParam(c *fiber.Ctx, paramName string) (uint, error) {
    // Parametre değerini al
    idParam := c.Params(paramName) // Örneğin, "id", "campaignId", "product_model_info_id"
    if idParam == "" {
        log.Println("Hata: ID parametresi eksik! Parametre adı:", paramName)
        return 0, errors.New("ID parametresi eksik")
    }

    // ID'yi uint'e çevir
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        log.Println("Hata: ID uint'e çevrilemedi -", err.Error())
        return 0, err
    }

    return uint(id), nil
}

// **Kullanıcı Doğrulama Fonksiyonu**
func ValidateUser(userID uint) error {
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		log.Println("Hata: Kullanıcı bulunamadı -", err.Error())
		return err
	}
	return nil
}

// CheckCampaignOwnership, kampanyanın kullanıcıya ait olup olmadığını kontrol eder.
func CheckCampaignOwnership(c *fiber.Ctx, campaignID uint, userID uint) (*models.Campaign, error) {
	var campaign models.Campaign
	if err := database.DB.First(&campaign, campaignID).Error; err != nil {
		log.Println("Hata: Kampanya bulunamadı -", err.Error())
		return nil, RespondWithError(c, fiber.StatusNotFound, "Kampanya bulunamadı")
	}

	if campaign.UserId != userID {
		log.Println("Hata: Kullanıcı yetkili değil -", userID)
		return nil, RespondWithError(c, fiber.StatusForbidden, "Bu işlemi yapmaya yetkiniz yok")
	}

	return &campaign, nil
}


// RespondWithError, hata mesajlarını standart bir şekilde döner
func RespondWithError(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
}




