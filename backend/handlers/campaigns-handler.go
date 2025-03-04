// package handlers

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"kfs-backend/models"
// 	"kfs-backend/services"
// 	"kfs-backend/utils"

// )

// // **Yeni Kampanya Oluştur**
// func CreateCampaign(c *fiber.Ctx) error {
// 	var req services.CampaignRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
// 	}

// 	if err := utils.ValidateUser(req.UserId); err != nil {
// 		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz UserId veya kullanıcı bulunamadı")
// 	}

// 	campaign := services.MapCampaignRequest(req, nil)

// 	createdCampaign, err := services.CreateCampaign(campaign)
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya oluşturulurken hata oluştu")
// 	}

// 	createdCampaign.User = models.User{}

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"message":  "Kampanya başarıyla oluşturuldu",
// 		"campaign": createdCampaign,
// 	})
// }

// // **Belirli Bir Kampanyayı Getir**
// func GetCampaignByID(c *fiber.Ctx) error {
// 	id, err := utils.GetIDParam(c, "id")
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
// 	}

// 	campaign, err := services.GetCampaignByID(id)
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusNotFound, "Kampanya bulunamadı")
// 	}

// 	return c.Status(fiber.StatusOK).JSON(campaign)
// }

// // **Tüm Kampanyaları Getir**
// func GetAllCampaigns(c *fiber.Ctx) error {
// 	campaigns, err := services.GetAllCampaigns()
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanyalar alınırken hata oluştu")
// 	}

// 	return c.Status(fiber.StatusOK).JSON(campaigns)
// }

// // **Kampanyayı Güncelle**
// func UpdateCampaign(c *fiber.Ctx) error {
// 	// Kullanıcı ID'sini güvenli bir şekilde al
// 	userIDInterface := c.Locals("userID")
// 	if userIDInterface == nil {
// 		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
// 	}

// 	currentUserID, ok := userIDInterface.(uint)
// 	if !ok {
// 		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
// 	}

// 	// Kampanya ID'sini al
// 	id, err := utils.GetIDParam(c, "id")
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
// 	}

// 	// Kullanıcının yetkisini kontrol et
// 	campaign, err := utils.CheckCampaignOwnership(c, id, currentUserID)
// 	if err != nil {
// 		return err
// 	}

// 	// JSON verisini parse et
// 	var req services.CampaignRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
// 	}

// 	updatedCampaign := services.MapCampaignRequest(req, campaign)
// 	savedCampaign, err := services.UpdateCampaign(id, updatedCampaign)
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya güncellenirken hata oluştu")
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message":  "Kampanya başarıyla güncellendi",
// 		"campaign": savedCampaign,
// 	})
// }

// // **Kampanyayı Sil**
// func DeleteCampaign(c *fiber.Ctx) error {
// 	// Kullanıcı ID'sini güvenli bir şekilde al
// 	userIDInterface := c.Locals("userID")
// 	if userIDInterface == nil {
// 		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
// 	}

// 	currentUserID, ok := userIDInterface.(uint)
// 	if !ok {
// 		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
// 	}

// 	// Kampanya ID'sini al
// 	id, err := utils.GetIDParam(c, "id")
// 	if err != nil {
// 		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
// 	}

// 	// Kampanyanın sahibini kontrol et
// 	_, err = utils.CheckCampaignOwnership(c, id, currentUserID)
// 	if err != nil {
// 		return err
// 	}

// 	// Kampanyayı sil
// 	if err := services.DeleteCampaign(id); err != nil {
// 		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya silinirken hata oluştu")
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kampanya başarıyla silindi"})
// }

package handlers

import (
	"kfs-backend/models"
	"kfs-backend/services"
	"kfs-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// **Yeni Kampanya Oluştur (S3 Desteği Eklenmiş)**
func CreateCampaign(c *fiber.Ctx) error {
	// JWT'den gelen userId'yi al
	userIDInterface := c.Locals("userId")
	if userIDInterface == nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
	}

	userID, ok := userIDInterface.(uint)
	if !ok {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
	}

	var req services.CampaignRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	if err := utils.ValidateUser(userID); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz UserId veya kullanıcı bulunamadı")
	}

	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	// Kampanya logoyu yükle (varsa)
	file, err := c.FormFile("campaign_logo") // Dosya yükleme
	if err == nil {                          // Eğer dosya varsa
		req.CampaignLogoKey, err = s3Service.UploadFile(file, "campaigns/logos")
		if err != nil {
			return utils.RespondWithError(c, fiber.StatusInternalServerError, "Dosya yükleme başarısız")
		}
	}

	campaign := services.MapCampaignRequest(req, nil, userID)
	createdCampaign, err := services.CreateCampaign(campaign)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya oluşturulurken hata oluştu")
	}

	createdCampaign.User = models.User{}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Kampanya başarıyla oluşturuldu",
		"campaign": createdCampaign,
	})
}

// **Belirli Bir Kampanyayı Getir**
func GetCampaignByID(c *fiber.Ctx) error {
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	campaign, err := services.GetCampaignByID(id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Kampanya bulunamadı")
	}

	return c.Status(fiber.StatusOK).JSON(campaign)
}

// **Tüm Kampanyaları Getir**
func GetAllCampaigns(c *fiber.Ctx) error {
	campaigns, err := services.GetAllCampaigns()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanyalar alınırken hata oluştu")
	}

	return c.Status(fiber.StatusOK).JSON(campaigns)
}

// **Kampanyayı Güncelle (S3 Desteği Eklenmiş)**
func UpdateCampaign(c *fiber.Ctx) error {
	userIDInterface := c.Locals("userID")
	if userIDInterface == nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
	}

	currentUserID, ok := userIDInterface.(uint)
	if !ok {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
	}

	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	campaign, err := utils.CheckCampaignOwnership(c, id, currentUserID)
	if err != nil {
		return err
	}

	var req services.CampaignRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	// Kampanya logosunu güncelle (varsa yeni dosya yüklenecek)
	file, err := c.FormFile("campaign_logo")
	if err == nil { // Eğer dosya varsa
		// Eski dosyayı sil
		if campaign.CampaignLogoKey != "" {
			s3Service.DeleteFile(campaign.CampaignLogoKey)
		}

		// Yeni dosyayı yükle
		req.CampaignLogoKey, err = s3Service.UploadFile(file, "campaigns/logos")
		if err != nil {
			return utils.RespondWithError(c, fiber.StatusInternalServerError, "Dosya yükleme başarısız")
		}
	}

	updatedCampaign := services.MapCampaignRequest(req, campaign, currentUserID)
	savedCampaign, err := services.UpdateCampaign(id, updatedCampaign)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya güncellenirken hata oluştu")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Kampanya başarıyla güncellendi",
		"campaign": savedCampaign,
	})
}

// **Kampanyayı Sil (S3 Desteği Eklenmiş)**
func DeleteCampaign(c *fiber.Ctx) error {
	userIDInterface := c.Locals("userID")
	if userIDInterface == nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
	}

	currentUserID, ok := userIDInterface.(uint)
	if !ok {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
	}

	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	campaign, err := utils.CheckCampaignOwnership(c, id, currentUserID)
	if err != nil {
		return err
	}

	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	// S3'ten kampanya logosunu sil
	if campaign.CampaignLogoKey != "" {
		s3Service.DeleteFile(campaign.CampaignLogoKey)
	}

	// Kampanyayı sil
	if err := services.DeleteCampaign(id); err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya silinirken hata oluştu")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kampanya başarıyla silindi"})
}