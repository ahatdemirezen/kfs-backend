package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kfs-backend/models"
	"kfs-backend/services"
	"kfs-backend/utils"
	
)

// **Yeni Kampanya Oluştur**
func CreateCampaign(c *fiber.Ctx) error {
	var req services.CampaignRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	if err := utils.ValidateUser(req.UserId); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz UserId veya kullanıcı bulunamadı")
	}

	campaign := services.MapCampaignRequest(req, nil)

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

// **Kampanyayı Güncelle**
func UpdateCampaign(c *fiber.Ctx) error {
	// Kullanıcı ID'sini güvenli bir şekilde al
	userIDInterface := c.Locals("userID")
	if userIDInterface == nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
	}

	currentUserID, ok := userIDInterface.(uint)
	if !ok {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
	}

	// Kampanya ID'sini al
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Kullanıcının yetkisini kontrol et
	campaign, err := utils.CheckCampaignOwnership(c, id, currentUserID)
	if err != nil {
		return err
	}

	// JSON verisini parse et
	var req services.CampaignRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	updatedCampaign := services.MapCampaignRequest(req, campaign)
	savedCampaign, err := services.UpdateCampaign(id, updatedCampaign)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya güncellenirken hata oluştu")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Kampanya başarıyla güncellendi",
		"campaign": savedCampaign,
	})
}


// **Kampanyayı Sil**
func DeleteCampaign(c *fiber.Ctx) error {
	// Kullanıcı ID'sini güvenli bir şekilde al
	userIDInterface := c.Locals("userID")
	if userIDInterface == nil {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Yetkisiz erişim")
	}
	
	currentUserID, ok := userIDInterface.(uint)
	if !ok {
		return utils.RespondWithError(c, fiber.StatusUnauthorized, "Geçersiz kullanıcı kimliği")
	}

	// Kampanya ID'sini al
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Kampanyanın sahibini kontrol et
	_, err = utils.CheckCampaignOwnership(c, id, currentUserID)
	if err != nil {
		return err
	}

	// Kampanyayı sil
	if err := services.DeleteCampaign(id); err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kampanya silinirken hata oluştu")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kampanya başarıyla silindi"})
}
