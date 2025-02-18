package services

import (

	"kfs-backend/database"
	"log"
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