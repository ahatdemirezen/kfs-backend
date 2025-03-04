package services

import (
	"fmt"
	"kfs-backend/database"
	"kfs-backend/models"
	"log"
	"reflect"
)

type VentureService[T any] interface {
	Create(req *T) (*T, error)
	GetByID(id uint) (*T, error)
	GetByField(field string, value interface{}) ([]T, error) // Yeni fonksiyon
	Update(id uint, req *T) (*T, error)
	Delete(id uint) error
}

type GenericVentureService[T any] struct{}

func (s *GenericVentureService[T]) Create(req *T) (*T, error) {
	if err := database.DB.Create(req).Error; err != nil {
		log.Println("Hata: Kayıt oluşturulamadı -", err.Error())
		return nil, err
	}
	return req, nil
}

func (s *GenericVentureService[T]) GetByID(id uint) (*T, error) {
	var record T
	if err := database.DB.First(&record, id).Error; err != nil {
		log.Println("Hata: Kayıt bulunamadı -", err.Error())
		return nil, err
	}
	return &record, nil
}

func (s *GenericVentureService[T]) GetByField(field string, value interface{}) ([]T, error) {
	var records []T
	if err := database.DB.Where(field+" = ?", value).Find(&records).Error; err != nil {
		log.Println("Hata: Kayıtlar bulunamadı -", err.Error())
		return nil, err
	}
	return records, nil
}

func (s *GenericVentureService[T]) Update(id uint, req *T) (*T, error) {
	var existing T
	if err := database.DB.First(&existing, id).Error; err != nil {
		log.Println("Hata: Güncellenecek kayıt bulunamadı -", err.Error())
		return nil, err
	}

	if err := database.DB.Model(&existing).Updates(req).Error; err != nil {
		log.Println("Hata: Kayıt güncellenemedi -", err.Error())
		return nil, err
	}

	return &existing, nil
}

func (s *GenericVentureService[T]) Delete(id uint) error {
	var record T
	if err := database.DB.Delete(&record, id).Error; err != nil {
		log.Println("Hata: Kayıt silinemedi -", err.Error())
		return err
	}
	return nil
}
func GetRecordByID(table string, id uint) (interface{}, error) {
	var record interface{}

	switch table {
	case "analysis_info":
		record = &models.AnalysisInfo{}
	case "campaigns":
		record = &models.Campaign{}
	case "patents":
		record = &models.Patent{}
	case "team_members":
		record = &models.TeamMember{}
	case "achievements":
		record = &models.Achievement{}
	case "permissions":
		record = &models.Permission{}
	case "product_model_infos":
		record = &models.ProductModelInfo{}
	case "other_product_topics":
		record = &models.OtherProductTopic{}
	case "market_infos":
		record = &models.MarketInfo{}
	case "other_market_topics":
		record = &models.OtherMarketTopic{}
	case "funding_infos":
		record = &models.FundingInfo{}
	case "after_funding_founder_partners":
		record = &models.AfterFundingFounderPartner{}
	case "visual_infos":
		record = &models.VisualInfo{}
	case "videos":
		record = &models.Videos{}
	case "other_photos":
		record = &models.OtherPhotos{}
	case "other_documents_infos":
		record = &models.OtherDocumentsInfo{}
	case "financial_documents":
		record = &models.FinancialDocuments{}
	default:
		return nil, fmt.Errorf("Bilinmeyen tablo: %s", table)
	}

	if err := database.DB.First(record, id).Error; err != nil {
		return nil, err
	}
	return record, nil
}

// Modelin belirtilen alanını günceller
func UpdateRecordField(record interface{}, field, value string) error {
	// Güncellenebilir alanların eşleşmesini sağla
	fieldMap := map[string]string{
		"swot_key":                      "SwotKey",
		"business_key":                  "BusinessKey",
		"investor_key":                  "InvestorKey",
		"campaign_logo_key":             "CampaignLogoKey",
		"document_key":                  "DocumentKey",
		"resume_key":                    "ResumeKey",
		"evaluation_report_key":         "EvaluationReportKey",
		"about_market_key":              "AboutMarketKey",
		"about_competition_key":         "AboutCompetitionKey",
		"target_summary_key":            "TargetSummaryKey",
		"commercialization_summary_key": "CommercializationSummaryKey",
		"photo_key":                     "PhotoKey",
		"about_product_key":             "AboutProductKey",
		"process_summary_key":           "ProcessSummaryKey",
		"about_process_key":             "AboutProcessKey",
		"about_side_key":                "AboutSideKey",
		"technical_analyses_key":        "TechnicalAnalysesKey",
		"arge_summary_key":              "ArgeSummaryKey",
		"previous_sales_key":            "PreviousSalesKey",
		"video_url":                     "VideoUrl",
		"showcase_photo_key":            "ShowcasePhotoKey",
	}
	structField, exists := fieldMap[field]
	if !exists {
		return fmt.Errorf("Güncellenecek alan bulunamadı: %s", field)
	}

	// Modelin ilgili alanına eriş
	val := reflect.ValueOf(record).Elem()
	fieldVal := val.FieldByName(structField)

	if !fieldVal.IsValid() || fieldVal.Kind() != reflect.String {
		return fmt.Errorf("Belirtilen alan güncellenemedi: %s", field)
	}

	// Yeni değeri ata
	fieldVal.SetString(value)

	// Veritabanına kaydet
	if err := database.DB.Save(record).Error; err != nil {
		return fmt.Errorf("Veritabanına kaydedilemedi: %s", err.Error())
	}

	return nil
}
