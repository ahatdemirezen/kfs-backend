package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// OtherMarketTopic oluştur
func CreateOtherMarketTopic(topic *models.OtherMarketTopic) error {
	db := database.DB

	// MarketInfo'nun varlığını kontrol et
	var marketInfo models.MarketInfo
	if err := db.First(&marketInfo, "market_info_id = ?", topic.MarketInfoId).Error; err != nil {
		return errors.New("related MarketInfo not found")
	}

	return db.Create(topic).Error
}

// ID'ye göre OtherMarketTopic getir
func GetOtherMarketTopicByID(id uint) (*models.OtherMarketTopic, error) {
	db := database.DB
	var topic models.OtherMarketTopic

	err := db.First(&topic, "topic_id = ?", id).Error
	if err != nil {
		return nil, errors.New("other market topic not found")
	}

	return &topic, nil
}

// Belirli bir MarketInfo'ya bağlı tüm OtherMarketTopic kayıtlarını getir
func GetOtherMarketTopicsByMarketInfoID(marketInfoId uint) ([]models.OtherMarketTopic, error) {
	db := database.DB
	var topics []models.OtherMarketTopic

	err := db.Where("market_info_id = ?", marketInfoId).Find(&topics).Error
	if err != nil {
		return nil, err
	}

	return topics, nil
}

// OtherMarketTopic güncelle
func UpdateOtherMarketTopic(topic *models.OtherMarketTopic) error {
	db := database.DB
	return db.Save(topic).Error
}

// OtherMarketTopic sil
func DeleteOtherMarketTopic(id uint) error {
	db := database.DB
	return db.Delete(&models.OtherMarketTopic{}, "topic_id = ?", id).Error
}
