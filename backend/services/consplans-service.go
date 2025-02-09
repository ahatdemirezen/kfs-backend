package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// ConsPlan oluştur
func CreateConsPlan(consPlan *models.ConsPlan) error {
	db := database.DB

	// AnalysisInfo'nun varlığını kontrol et
	var analysisInfo models.AnalysisInfo
	if err := db.First(&analysisInfo, "analysis_info_id = ?", consPlan.AnalysisInfoId).Error; err != nil {
		return errors.New("related AnalysisInfo not found")
	}

	return db.Create(consPlan).Error
}

// ID'ye göre ConsPlan getir
func GetConsPlanByID(id uint) (*models.ConsPlan, error) {
	db := database.DB
	var consPlan models.ConsPlan

	err := db.First(&consPlan, "con_plan_id = ?", id).Error
	if err != nil {
		return nil, errors.New("cons plan not found")
	}

	return &consPlan, nil
}

// Belirli bir AnalysisInfo'ya bağlı tüm ConsPlan kayıtlarını getir
func GetConsPlansByAnalysisInfoID(analysisInfoId uint) ([]models.ConsPlan, error) {
	db := database.DB
	var consPlans []models.ConsPlan

	err := db.Where("analysis_info_id = ?", analysisInfoId).Find(&consPlans).Error
	if err != nil {
		return nil, err
	}

	return consPlans, nil
}

// ConsPlan güncelle
func UpdateConsPlan(consPlan *models.ConsPlan) error {
	db := database.DB
	return db.Save(consPlan).Error
}

// ConsPlan sil
func DeleteConsPlan(id uint) error {
	db := database.DB
	return db.Delete(&models.ConsPlan{}, "con_plan_id = ?", id).Error
}

// Birden fazla ConsPlan oluştur
func CreateMultipleConsPlans(consPlans []models.ConsPlan) error {
	db := database.DB

	for _, consPlan := range consPlans {
		// AnalysisInfo'nun varlığını kontrol et
		var analysisInfo models.AnalysisInfo
		if err := db.First(&analysisInfo, "analysis_info_id = ?", consPlan.AnalysisInfoId).Error; err != nil {
			return errors.New("related AnalysisInfo not found for cons plan: " + consPlan.ConPlan)
		}

		// ConsPlan kaydını oluştur
		if err := db.Create(&consPlan).Error; err != nil {
			return err
		}
	}

	return nil
}
