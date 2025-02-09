package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// MarketInfo oluştur
func CreateMarketInfo(marketInfo *models.MarketInfo) error {
	db := database.DB

	// Campaign'in varlığını kontrol et
	var campaign models.Campaign
	if err := db.First(&campaign, "campaign_id = ?", marketInfo.CampaignId).Error; err != nil {
		return errors.New("related Campaign not found")
	}

	return db.Create(marketInfo).Error
}

// ID'ye göre MarketInfo getir
func GetMarketInfoByID(id uint) (*models.MarketInfo, error) {
	db := database.DB
	var marketInfo models.MarketInfo

	err := db.First(&marketInfo, "market_info_id = ?", id).Error
	if err != nil {
		return nil, errors.New("market info not found")
	}

	return &marketInfo, nil
}

// Belirli bir Campaign'e bağlı tüm MarketInfo kayıtlarını getir
func GetMarketInfosByCampaignID(campaignId uint) ([]models.MarketInfo, error) {
	db := database.DB
	var marketInfos []models.MarketInfo

	err := db.Where("campaign_id = ?", campaignId).Find(&marketInfos).Error
	if err != nil {
		return nil, err
	}

	return marketInfos, nil
}

// MarketInfo güncelle
func UpdateMarketInfo(marketInfo *models.MarketInfo) error {
	db := database.DB
	return db.Save(marketInfo).Error
}

// MarketInfo sil
func DeleteMarketInfo(id uint) error {
	db := database.DB
	return db.Delete(&models.MarketInfo{}, "market_info_id = ?", id).Error
}
