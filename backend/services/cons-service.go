package services

import (
	"errors"
	"kfs-backend/database"
	"kfs-backend/models"
)

// Cons oluştur
func CreateCon(con *models.Cons) error {
	db := database.DB

	// AnalysisInfo'nun varlığını kontrol et
	var analysisInfo models.AnalysisInfo
	if err := db.First(&analysisInfo, "analysis_info_id = ?", con.AnalysisInfoId).Error; err != nil {
		return errors.New("related AnalysisInfo not found")
	}

	return db.Create(con).Error
}

// ID'ye göre Cons getir
func GetConByID(id uint) (*models.Cons, error) {
	db := database.DB
	var con models.Cons

	// Veritabanında ara
	err := db.First(&con, "con_id = ?", id).Error
	if err != nil {
		return nil, errors.New("con not found")
	}

	return &con, nil
}

// Belirli bir AnalysisInfo'ya bağlı tüm Cons kayıtlarını getir
func GetConsByAnalysisInfoID(analysisInfoId uint) ([]models.Cons, error) {
	db := database.DB
	var cons []models.Cons

	err := db.Where("analysis_info_id = ?", analysisInfoId).Find(&cons).Error
	if err != nil {
		return nil, err
	}

	return cons, nil
}

// Cons güncelle
func UpdateCon(con *models.Cons) error {
	db := database.DB
	return db.Save(con).Error
}

// Cons sil
func DeleteCon(id uint) error {
	db := database.DB
	return db.Delete(&models.Cons{}, "con_id = ?", id).Error
}

// Birden fazla Cons oluştur
func CreateMultipleCons(cons []models.Cons) error {
	db := database.DB

	for _, con := range cons {
		// AnalysisInfo'nun varlığını kontrol et
		var analysisInfo models.AnalysisInfo
		if err := db.First(&analysisInfo, "analysis_info_id = ?", con.AnalysisInfoId).Error; err != nil {
			return errors.New("related AnalysisInfo not found for con: " + con.Con)
		}

		// Cons kaydını oluştur
		if err := db.Create(&con).Error; err != nil {
			return err
		}
	}

	return nil
}
