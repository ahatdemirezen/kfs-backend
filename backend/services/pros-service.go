package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// Pros oluştur
func CreatePro(pro *models.Pros) error {
	db := database.DB

	// AnalysisInfo'nun var olup olmadığını kontrol et
	var analysisInfo models.AnalysisInfo
	if err := db.First(&analysisInfo, "analysis_info_id = ?", pro.AnalysisInfoId).Error; err != nil {
		return errors.New("related AnalysisInfo not found")
	}

	return db.Create(pro).Error
}

// ID'ye göre Pros getir
func GetProByID(id uint) (*models.Pros, error) {
	db := database.DB
	var pro models.Pros

	// Veritabanında ara
	err := db.First(&pro, "pro_id = ?", id).Error
	if err != nil {
		return nil, errors.New("pro not found")
	}

	return &pro, nil
}

// Belirli bir AnalysisInfo'ya bağlı tüm Pros kayıtlarını getir
func GetProsByAnalysisInfoID(analysisInfoId uint) ([]models.Pros, error) {
	db := database.DB
	var pros []models.Pros

	// Veritabanında ilgili tüm Pros kayıtlarını al
	err := db.Where("analysis_info_id = ?", analysisInfoId).Find(&pros).Error
	if err != nil {
		return nil, err
	}

	return pros, nil
}

// Pros güncelle
func UpdatePro(pro *models.Pros) error {
	db := database.DB
	return db.Save(pro).Error
}

// Pros sil
func DeletePro(id uint) error {
	db := database.DB
	return db.Delete(&models.Pros{}, "pro_id = ?", id).Error
}

// Birden fazla Pros oluştur
func CreateMultiplePros(pros []models.Pros) error {
	db := database.DB

	for _, pro := range pros {
		// AnalysisInfo'nun var olup olmadığını kontrol et
		var analysisInfo models.AnalysisInfo
		if err := db.First(&analysisInfo, "analysis_info_id = ?", pro.AnalysisInfoId).Error; err != nil {
			return errors.New("related AnalysisInfo not found for pro: " + pro.Pro)
		}

		// Pros'u oluştur
		if err := db.Create(&pro).Error; err != nil {
			return err
		}
	}

	return nil
}
