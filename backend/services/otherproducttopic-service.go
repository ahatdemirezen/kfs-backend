package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// OtherProductTopic oluştur
func CreateOtherProductTopic(topic *models.OtherProductTopic) (*models.OtherProductTopic, error) {
	db := database.DB

	// ProductModelInfo'nun varlığını kontrol et
	var productModelInfo models.ProductModelInfo
	if err := db.First(&productModelInfo, "product_model_info_id = ?", topic.ProductModelInfoId).Error; err != nil {
		return nil, errors.New("related ProductModelInfo not found")
	}

	// Yeni OtherProductTopic oluştur
	if err := db.Create(topic).Error; err != nil {
		return nil, err
	}

	// Oluşturulan kaydı ProductModelInfo ve Campaign ile preload ederek geri dön
	if err := db.Preload("ProductModelInfo.Campaign").First(topic, "topic_id = ?", topic.TopicId).Error; err != nil {
		return nil, err
	}

	return topic, nil
}

// ID'ye göre OtherProductTopic getir
func GetOtherProductTopicByID(id uint) (*models.OtherProductTopic, error) {
	db := database.DB
	var topic models.OtherProductTopic

	// ProductModelInfo ve bağlı Campaign'i preload edin
	err := db.Preload("ProductModelInfo.Campaign").First(&topic, "topic_id = ?", id).Error
	if err != nil {
		return nil, errors.New("other product topic not found")
	}

	return &topic, nil
}

// Belirli bir ProductModelInfo'ya bağlı tüm OtherProductTopic kayıtlarını getir
func GetOtherProductTopicsByProductModelInfoID(productModelInfoId uint) ([]models.OtherProductTopic, error) {
	db := database.DB
	var topics []models.OtherProductTopic

	// Preload ile ilişkili ProductModelInfo'yu da getir
	err := db.Preload("ProductModelInfo").Where("product_model_info_id = ?", productModelInfoId).Find(&topics).Error
	if err != nil {
		return nil, err
	}

	return topics, nil
}

// OtherProductTopic güncelle
func UpdateOtherProductTopic(topic *models.OtherProductTopic) error {
	db := database.DB
	return db.Save(topic).Error
}

// OtherProductTopic sil
func DeleteOtherProductTopic(id uint) error {
	db := database.DB
	return db.Delete(&models.OtherProductTopic{}, "topic_id = ?", id).Error
}
