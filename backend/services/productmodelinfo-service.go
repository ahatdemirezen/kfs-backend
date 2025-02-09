package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// ProductModelInfo oluştur
func CreateProductModelInfo(info *models.ProductModelInfo) error {
	db := database.DB

	// Campaign'in varlığını kontrol et
	var campaign models.Campaign
	if err := db.First(&campaign, "campaign_id = ?", info.CampaignId).Error; err != nil {
		return errors.New("related Campaign not found")
	}

	// ProductModelInfo kaydını oluştur
	if err := db.Create(info).Error; err != nil {
		return err
	}

	// Campaign bilgilerini preload ederek doldur
	if err := db.Preload("Campaign").First(info, "product_model_info_id = ?", info.ProductModelInfoId).Error; err != nil {
		return err
	}

	return nil
}

// ID'ye göre ProductModelInfo getir
func GetProductModelInfoByID(id uint) (*models.ProductModelInfo, error) {
	db := database.DB
	var info models.ProductModelInfo

	// Preload kullanılarak Campaign bilgisi yükleniyor
	err := db.Preload("Campaign").First(&info, "product_model_info_id = ?", id).Error
	if err != nil {
		return nil, errors.New("product model info not found")
	}

	return &info, nil
}

// Belirli bir Campaign'e bağlı tüm ProductModelInfo kayıtlarını getir
func GetProductModelInfosByCampaignID(campaignId uint) ([]models.ProductModelInfo, error) {
	db := database.DB
	var infos []models.ProductModelInfo

	// Preload ile Campaign bilgisi ekleniyor
	err := db.Preload("Campaign").Where("campaign_id = ?", campaignId).Find(&infos).Error
	if err != nil {
		return nil, err
	}

	return infos, nil
}

// ProductModelInfo güncelle
func UpdateProductModelInfo(info *models.ProductModelInfo) error {
	db := database.DB
	return db.Save(info).Error
}

// ProductModelInfo sil
func DeleteProductModelInfo(id uint) error {
	db := database.DB
	return db.Delete(&models.ProductModelInfo{}, "product_model_info_id = ?", id).Error
}
