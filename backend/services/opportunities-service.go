package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// Opportunity oluştur
func CreateOpportunity(opportunity *models.Opportunity) error {
	db := database.DB

	// AnalysisInfo'nun var olup olmadığını kontrol et
	var analysisInfo models.AnalysisInfo
	if err := db.First(&analysisInfo, "analysis_info_id = ?", opportunity.AnalysisInfoId).Error; err != nil {
		return errors.New("related AnalysisInfo not found")
	}

	return db.Create(opportunity).Error
}

// ID'ye göre Opportunity getir
func GetOpportunityByID(id uint) (*models.Opportunity, error) {
	db := database.DB
	var opportunity models.Opportunity

	err := db.First(&opportunity, "opportunity_id = ?", id).Error
	if err != nil {
		return nil, errors.New("opportunity not found")
	}

	return &opportunity, nil
}

// Belirli bir AnalysisInfo'ya bağlı tüm Opportunities kayıtlarını getir
func GetOpportunitiesByAnalysisInfoID(analysisInfoId uint) ([]models.Opportunity, error) {
	db := database.DB
	var opportunities []models.Opportunity

	err := db.Where("analysis_info_id = ?", analysisInfoId).Find(&opportunities).Error
	if err != nil {
		return nil, err
	}

	return opportunities, nil
}

// Opportunity güncelle
func UpdateOpportunity(opportunity *models.Opportunity) error {
	db := database.DB
	return db.Save(opportunity).Error
}

// Opportunity sil
func DeleteOpportunity(id uint) error {
	db := database.DB
	return db.Delete(&models.Opportunity{}, "opportunity_id = ?", id).Error
}

// Birden fazla Opportunity oluştur
func CreateMultipleOpportunities(opportunities []models.Opportunity) error {
	db := database.DB

	for _, opportunity := range opportunities {
		// AnalysisInfo'nun varlığını kontrol et
		var analysisInfo models.AnalysisInfo
		if err := db.First(&analysisInfo, "analysis_info_id = ?", opportunity.AnalysisInfoId).Error; err != nil {
			return errors.New("related AnalysisInfo not found for opportunity: " + opportunity.Opportunity)
		}

		// Opportunity kaydını oluştur
		if err := db.Create(&opportunity).Error; err != nil {
			return err
		}
	}

	return nil
}
