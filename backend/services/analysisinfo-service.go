package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// AnalysisInfo oluştur
func CreateAnalysisInfo(analysisInfo *models.AnalysisInfo) error {
	db := database.DB
	return db.Create(analysisInfo).Error
}

// ID'ye göre AnalysisInfo getir
func GetAnalysisInfoByID(id uint) (*models.AnalysisInfo, error) {
	db := database.DB
	var analysisInfo models.AnalysisInfo

	// Veritabanında ara
	err := db.First(&analysisInfo, "analysis_info_id = ?", id).Error
	if err != nil {
		return nil, errors.New("analysis info not found")
	}

	return &analysisInfo, nil
}

// AnalysisInfo güncelle
func UpdateAnalysisInfo(analysisInfo *models.AnalysisInfo) error {
	db := database.DB
	return db.Save(analysisInfo).Error
}

// AnalysisInfo sil
func DeleteAnalysisInfo(id uint) error {
	db := database.DB
	return db.Delete(&models.AnalysisInfo{}, "analysis_info_id = ?", id).Error
}
