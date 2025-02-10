package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// MarketInfo oluştur
func CreateMarketInfo(marketInfo *models.MarketInfo) (*models.MarketInfo, error) {
	db := database.DB

	// Campaign'in varlığını kontrol et
	var campaign models.Campaign
	if err := db.First(&campaign, "campaign_id = ?", marketInfo.CampaignId).Error; err != nil {
		return nil, errors.New("related Campaign not found")
	}

	// MarketInfo oluştur
	if err := db.Create(marketInfo).Error; err != nil {
		return nil, err
	}

	// Preload işlemi ile ilişkileri doldur
	if err := db.Preload("Campaign").Find(marketInfo).Error; err != nil {
		return nil, err
	}

	return marketInfo, nil
}

// ID'ye göre MarketInfo getir
func GetMarketInfoByID(id uint) (*models.MarketInfo, error) {
	db := database.DB
	var marketInfo models.MarketInfo

	// Preload işlemi ile Campaign ilişkisini getir
	err := db.Preload("Campaign").First(&marketInfo, "market_info_id = ?", id).Error
	if err != nil {
		return nil, errors.New("market info not found")
	}

	return &marketInfo, nil
}

// Belirli bir Campaign'e bağlı tüm MarketInfo kayıtlarını getir
func GetMarketInfosByCampaignID(campaignId uint) ([]models.MarketInfo, error) {
	db := database.DB
	var marketInfos []models.MarketInfo

	// Campaign ile ilişkili MarketInfo'ları getir ve Campaign ilişkisini doldur
	err := db.Where("campaign_id = ?", campaignId).Preload("Campaign").Find(&marketInfos).Error
	if err != nil {
		return nil, err
	}

	return marketInfos, nil
}

// MarketInfo güncelle
func UpdateMarketInfo(marketInfo *models.MarketInfo) (*models.MarketInfo, error) {
	db := database.DB

	// MarketInfo güncelle
	if err := db.Save(marketInfo).Error; err != nil {
		return nil, err
	}

	// Güncellenmiş kaydı Campaign ilişkisiyle birlikte getir
	if err := db.Preload("Campaign").Find(marketInfo, "market_info_id = ?", marketInfo.MarketInfoId).Error; err != nil {
		return nil, err
	}

	return marketInfo, nil
}

// MarketInfo sil
func DeleteMarketInfo(id uint) error {
	db := database.DB
	return db.Delete(&models.MarketInfo{}, "market_info_id = ?", id).Error
}
