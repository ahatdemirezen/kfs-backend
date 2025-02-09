package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// ThreatPlan oluştur
func CreateThreatPlan(threatPlan *models.ThreatPlan) error {
	db := database.DB

	// AnalysisInfo'nun varlığını kontrol et
	var analysisInfo models.AnalysisInfo
	if err := db.First(&analysisInfo, "analysis_info_id = ?", threatPlan.AnalysisInfoId).Error; err != nil {
		return errors.New("related AnalysisInfo not found")
	}

	return db.Create(threatPlan).Error
}

// ID'ye göre ThreatPlan getir
func GetThreatPlanByID(id uint) (*models.ThreatPlan, error) {
	db := database.DB
	var threatPlan models.ThreatPlan

	err := db.First(&threatPlan, "threat_plan_id = ?", id).Error
	if err != nil {
		return nil, errors.New("threat plan not found")
	}

	return &threatPlan, nil
}

// Belirli bir AnalysisInfo'ya bağlı tüm ThreatPlan kayıtlarını getir
func GetThreatPlansByAnalysisInfoID(analysisInfoId uint) ([]models.ThreatPlan, error) {
	db := database.DB
	var threatPlans []models.ThreatPlan

	err := db.Where("analysis_info_id = ?", analysisInfoId).Find(&threatPlans).Error
	if err != nil {
		return nil, err
	}

	return threatPlans, nil
}

// ThreatPlan güncelle
func UpdateThreatPlan(threatPlan *models.ThreatPlan) error {
	db := database.DB
	return db.Save(threatPlan).Error
}

// ThreatPlan sil
func DeleteThreatPlan(id uint) error {
	db := database.DB
	return db.Delete(&models.ThreatPlan{}, "threat_plan_id = ?", id).Error
}

// Birden fazla ThreatPlan oluştur
func CreateMultipleThreatPlans(threatPlans []models.ThreatPlan) error {
	db := database.DB

	for _, threatPlan := range threatPlans {
		// AnalysisInfo'nun varlığını kontrol et
		var analysisInfo models.AnalysisInfo
		if err := db.First(&analysisInfo, "analysis_info_id = ?", threatPlan.AnalysisInfoId).Error; err != nil {
			return errors.New("related AnalysisInfo not found for threat plan: " + threatPlan.ThreatPlan)
		}

		// ThreatPlan kaydını oluştur
		if err := db.Create(&threatPlan).Error; err != nil {
			return err
		}
	}

	return nil
}
