// package services

// import (
// 	"log"
// 	"kfs-backend/database"
// 	"kfs-backend/models"
// 	"time"
// )

// // **CampaignRequest Yapısı (`handlers` yerine `services` içinde)**
// type CampaignRequest struct {
// 	UserId          uint   `json:"user_id" validate:"required"`
// 	CampaignStatus  string `json:"campaign_status"`
// 	CampaignCode    string `json:"campaign_code"`
// 	CampaignLogoKey string `json:"campaign_logo_key"`
// 	VentureName     string `json:"venture_name"`
// 	CampaignName    string `json:"campaign_name" validate:"required"`
// 	Description     string `json:"description"`
// 	AboutProject    string `json:"about_project"`
// 	Summary         string `json:"summary"`
// 	VenturePurpose  string `json:"venture_purpose"`
// 	VenturePhase    string `json:"venture_phase"`
// }

// // **Yeni Kampanya Oluştur**
// func CreateCampaign(campaign *models.Campaign) (*models.Campaign, error) {
// 	if err := database.DB.Create(&campaign).Error; err != nil {
// 		log.Println("Hata: Kampanya oluşturulamadı -", err.Error())
// 		return nil, err
// 	}
// 	return campaign, nil
// }

// // **ID ile Kampanya Getir**
// func GetCampaignByID(id uint) (*models.Campaign, error) {
// 	var campaign models.Campaign
// 	if err := database.DB.First(&campaign, id).Error; err != nil {
// 		log.Println("Hata: Kampanya bulunamadı -", err.Error())
// 		return nil, err
// 	}
// 	campaign.User = models.User{} // **Hassas bilgileri temizle**
// 	return &campaign, nil
// }

// // **Tüm Kampanyaları Getir**
// func GetAllCampaigns() ([]models.Campaign, error) {
// 	var campaigns []models.Campaign
// 	if err := database.DB.Find(&campaigns).Error; err != nil {
// 		log.Println("Hata: Kampanyalar alınamadı -", err.Error())
// 		return nil, err
// 	}

// 	// **User bilgilerini temizle**
// 	for i := range campaigns {
// 		campaigns[i].User = models.User{}
// 	}

// 	return campaigns, nil
// }

// // **Kampanyayı Güncelle**
// func UpdateCampaign(id uint, campaign *models.Campaign) (*models.Campaign, error) {
// 	var existingCampaign models.Campaign
// 	if err := database.DB.First(&existingCampaign, id).Error; err != nil {
// 		log.Println("Hata: Güncellenecek kampanya bulunamadı -", err.Error())
// 		return nil, err
// 	}

// 	// **Yalnızca değişen alanları güncelle**
// 	if err := database.DB.Model(&existingCampaign).Updates(campaign).Error; err != nil {
// 		log.Println("Hata: Kampanya güncellenemedi -", err.Error())
// 		return nil, err
// 	}

// 	return &existingCampaign, nil
// }


// // **Kampanyayı Sil**
// func DeleteCampaign(id uint) error {
// 	if err := database.DB.Delete(&models.Campaign{}, id).Error; err != nil {
// 		log.Println("Hata: Kampanya silinemedi -", err.Error())
// 		return err
// 	}
// 	return nil
// }



// // **CampaignRequest'i Campaign Modeline Map'leme**
// func MapCampaignRequest(req CampaignRequest, existing *models.Campaign) *models.Campaign {
// 	if existing == nil {
// 		existing = &models.Campaign{CreatedAt: time.Now()}
// 	}

// 	existing.UserId = req.UserId
// 	existing.CampaignStatus = req.CampaignStatus
// 	existing.CampaignCode = req.CampaignCode
// 	existing.CampaignLogoKey = req.CampaignLogoKey
// 	existing.VentureName = req.VentureName
// 	existing.CampaignName = req.CampaignName
// 	existing.Description = req.Description
// 	existing.AboutProject = req.AboutProject
// 	existing.Summary = req.Summary
// 	existing.VenturePurpose = req.VenturePurpose
// 	existing.VenturePhase = req.VenturePhase
// 	existing.UpdatedAt = time.Now()

// 	return existing
// }





package services

import (
	"log"
	"kfs-backend/database"
	"kfs-backend/models"
	"time"
)

// **CampaignRequest Yapısı**
type CampaignRequest struct {
	CampaignStatus  string `json:"campaign_status"`
	CampaignCode    string `json:"campaign_code"`
	CampaignLogoKey string `json:"campaign_logo_key"`
	VentureName     string `json:"venture_name"`
	CampaignName    string `json:"campaign_name" validate:"required"`
	Description     string `json:"description"`
	AboutProject    string `json:"about_project"`
	Summary         string `json:"summary"`
	VenturePurpose  string `json:"venture_purpose"`
	VenturePhase    string `json:"venture_phase"`
}

// **Yeni Kampanya Oluştur**
func CreateCampaign(campaign *models.Campaign) (*models.Campaign, error) {
	if err := database.DB.Create(&campaign).Error; err != nil {
		log.Println("Hata: Kampanya oluşturulamadı -", err.Error())
		return nil, err
	}
	return campaign, nil
}

// **ID ile Kampanya Getir**
func GetCampaignByID(id uint) (*models.Campaign, error) {
	var campaign models.Campaign
	if err := database.DB.First(&campaign, id).Error; err != nil {
		log.Println("Hata: Kampanya bulunamadı -", err.Error())
		return nil, err
	}
	campaign.User = models.User{} // **Hassas bilgileri temizle**
	return &campaign, nil
}

// **Tüm Kampanyaları Getir**
func GetAllCampaigns() ([]models.Campaign, error) {
	var campaigns []models.Campaign
	if err := database.DB.Find(&campaigns).Error; err != nil {
		log.Println("Hata: Kampanyalar alınamadı -", err.Error())
		return nil, err
	}

	// **User bilgilerini temizle**
	for i := range campaigns {
		campaigns[i].User = models.User{}
	}

	return campaigns, nil
}

// **Kampanyayı Güncelle**
func UpdateCampaign(id uint, campaign *models.Campaign) (*models.Campaign, error) {
	var existingCampaign models.Campaign
	if err := database.DB.First(&existingCampaign, id).Error; err != nil {
		log.Println("Hata: Güncellenecek kampanya bulunamadı -", err.Error())
		return nil, err
	}

	// **Yalnızca değişen alanları güncelle**
	if err := database.DB.Model(&existingCampaign).Updates(campaign).Error; err != nil {
		log.Println("Hata: Kampanya güncellenemedi -", err.Error())
		return nil, err
	}

	return &existingCampaign, nil
}

// **Kampanyayı Sil**
func DeleteCampaign(id uint) error {
	if err := database.DB.Delete(&models.Campaign{}, id).Error; err != nil {
		log.Println("Hata: Kampanya silinemedi -", err.Error())
		return err
	}
	return nil
}

// **CampaignRequest'i Campaign Modeline Map'leme**
func MapCampaignRequest(req CampaignRequest, existing *models.Campaign, userID uint) *models.Campaign {
	if existing == nil {
		existing = &models.Campaign{CreatedAt: time.Now()}
		existing.UserId = userID // **UserId doğrudan JWT’den alınır**
	}

	existing.CampaignStatus = req.CampaignStatus
	existing.CampaignCode = req.CampaignCode
	existing.CampaignLogoKey = req.CampaignLogoKey
	existing.VentureName = req.VentureName
	existing.CampaignName = req.CampaignName
	existing.Description = req.Description
	existing.AboutProject = req.AboutProject
	existing.Summary = req.Summary
	existing.VenturePurpose = req.VenturePurpose
	existing.VenturePhase = req.VenturePhase
	existing.UpdatedAt = time.Now()

	return existing
}
