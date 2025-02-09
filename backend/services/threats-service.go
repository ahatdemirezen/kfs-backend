package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// Threat oluştur
func CreateThreat(threat *models.Threat) error {
	db := database.DB

	// AnalysisInfo'nun varlığını kontrol et
	var analysisInfo models.AnalysisInfo
	if err := db.First(&analysisInfo, "analysis_info_id = ?", threat.AnalysisInfoId).Error; err != nil {
		return errors.New("related AnalysisInfo not found")
	}

	return db.Create(threat).Error
}

// ID'ye göre Threat getir
func GetThreatByID(id uint) (*models.Threat, error) {
	db := database.DB
	var threat models.Threat

	err := db.First(&threat, "threat_id = ?", id).Error
	if err != nil {
		return nil, errors.New("threat not found")
	}

	return &threat, nil
}

// Belirli bir AnalysisInfo'ya bağlı tüm Threat kayıtlarını getir
func GetThreatsByAnalysisInfoID(analysisInfoId uint) ([]models.Threat, error) {
	db := database.DB
	var threats []models.Threat

	err := db.Where("analysis_info_id = ?", analysisInfoId).Find(&threats).Error
	if err != nil {
		return nil, err
	}

	return threats, nil
}

// Threat güncelle
func UpdateThreat(threat *models.Threat) error {
	db := database.DB
	return db.Save(threat).Error
}

// Threat sil
func DeleteThreat(id uint) error {
	db := database.DB
	return db.Delete(&models.Threat{}, "threat_id = ?", id).Error
}

// Birden fazla Threat oluştur
func CreateMultipleThreats(threats []models.Threat) error {
	db := database.DB

	for _, threat := range threats {
		// AnalysisInfo'nun varlığını kontrol et
		var analysisInfo models.AnalysisInfo
		if err := db.First(&analysisInfo, "analysis_info_id = ?", threat.AnalysisInfoId).Error; err != nil {
			return errors.New("related AnalysisInfo not found for threat: " + threat.Threat)
		}

		// Threat kaydını oluştur
		if err := db.Create(&threat).Error; err != nil {
			return err
		}
	}

	return nil
}
