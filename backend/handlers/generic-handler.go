package handlers

import (
	"fmt"
	"kfs-backend/services"
	"kfs-backend/utils"

	"github.com/gofiber/fiber/v2"
)

// Yeni bir girişimi oluşturur.
func CreateVenture[T any, R any](c *fiber.Ctx, service services.VentureService[T], req R) error {
	// JSON verisini parse et
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	var fileKey string
	// Dosya yükleme işlemi (varsa)
	file, err := c.FormFile("file")
	if err == nil { // Eğer dosya varsa yükle
		fileKey, err = s3Service.UploadFile(file, "ventures/files")
		if err != nil {
			return utils.RespondWithError(c, fiber.StatusInternalServerError, "Dosya yükleme başarısız")
		}
	}

	// Modeli oluştur ve dosya anahtarını ekleyerek kaydet
	model := utils.ConvertRequestToModel[R, T](req, fileKey)

	// Servis ile kaydı oluştur
	createdRecord, err := service.Create(&model)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kayıt oluşturulurken hata oluştu")
	}

	// Başarıyla oluşturulan kaydı döndür
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kayıt başarıyla oluşturuldu",
		"data":    createdRecord,
	})
}

// ID ile girişim bilgisini getirir.
func GetVentureByID[T any](c *fiber.Ctx, service services.VentureService[T]) error {
	// ID parametresini al
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Servis üzerinden kaydı getir
	record, err := service.GetByID(id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Kayıt bulunamadı")
	}

	// Kaydı döndür
	return c.Status(fiber.StatusOK).JSON(record)
}

// Belirli bir alana göre girişim kayıtlarını getirir.
func GetVenturesByField[T any](c *fiber.Ctx, service services.VentureService[T], field string) error {
	// Parametre değerini al
	value, err := utils.GetIDParam(c, field)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz parametre değeri")
	}

	// Servis üzerinden kayıtları getir
	records, err := service.GetByField(field, value)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Kayıtlar bulunamadı")
	}

	// Kayıtları döndür
	return c.Status(fiber.StatusOK).JSON(records)
}

// Girişim bilgilerini günceller.
func UpdateVenture[T any, R any](c *fiber.Ctx, service services.VentureService[T], req R) error {
	// ID parametresini al
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// JSON verisini parse et
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	// Mevcut kaydı al
	existingRecord, err := service.GetByID(id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Güncellenecek kayıt bulunamadı")
	}

	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	var fileKey string
	// Dosya yükleme işlemi (varsa yeni dosya yüklenirse, eskisini sil)
	file, err := c.FormFile("file")
	if err == nil {
		// Önceki dosyayı sil (varsa)
		if existingFileKey, ok := utils.GetFileKey(existingRecord); ok && existingFileKey != "" {
			s3Service.DeleteFile(existingFileKey)
		}

		// Yeni dosyayı yükle
		fileKey, err = s3Service.UploadFile(file, "ventures/files")
		if err != nil {
			return utils.RespondWithError(c, fiber.StatusInternalServerError, "Dosya yükleme başarısız")
		}
	} else {
		// Yeni dosya yoksa, mevcut dosya anahtarını koru
		if existingFileKey, ok := utils.GetFileKey(existingRecord); ok {
			fileKey = existingFileKey
		}
	}

	// Modeli oluştur ve dosya anahtarını güncelleyerek kaydet
	model := utils.ConvertRequestToModel[R, T](req, fileKey)

	// Servis üzerinden kaydı güncelle
	updatedRecord, err := service.Update(id, &model)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kayıt güncellenirken hata oluştu")
	}

	// Güncellenen kaydı döndür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Kayıt başarıyla güncellendi",
		"data":    updatedRecord,
	})
}

// ID ile girişim kaydını siler.
func DeleteVenture[T any](c *fiber.Ctx, service services.VentureService[T]) error {
	// ID parametresini al
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Mevcut kaydı al
	existingRecord, err := service.GetByID(id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Silinecek kayıt bulunamadı")
	}

	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	// Kayıtla ilişkili dosyayı sil (varsa)
	if existingFileKey, ok := utils.GetFileKey(existingRecord); ok && existingFileKey != "" {
		s3Service.DeleteFile(existingFileKey)
	}

	// Servis üzerinden kaydı sil
	if err := service.Delete(id); err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kayıt silinirken hata oluştu")
	}

	// Başarı mesajını döndür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kayıt başarıyla silindi"})
}

// Generic Dosya Yükleme Handler
func UploadFileAndUpdate(c *fiber.Ctx, service services.VentureService[any]) error {
	// S3 Servisini başlat
	s3Service, err := services.NewS3Service()
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "S3 servisi başlatılamadı")
	}

	// Dosyayı al
	file, err := c.FormFile("file")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Dosya yüklenemedi, lütfen bir dosya seçin")
	}

	// Hangi tabloya (sekme) yükleneceğini belirten parametre
	table := c.FormValue("table")
	if table == "" {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz table parametresi")
	}

	// Hangi alana (field) ekleneceğini belirten parametre
	field := c.FormValue("field")
	if field == "" {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz field parametresi")
	}

	// Güncellenmesi gereken kaydın ID'si
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Hangi klasöre yükleneceğini belirle (Örneğin: campaigns/files, patents/files)
	folder := fmt.Sprintf("%s/files", table)

	// Dosyayı S3'e yükle
	fileKey, err := s3Service.UploadFile(file, folder)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Dosya yükleme başarısız")
	}

	// Dinamik model getirme (hangi tablo güncellenecekse onun modelini buluyoruz)
	record, err := services.GetRecordByID(table, id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, fmt.Sprintf("%s kaydı bulunamadı", table))
	}

	// Güncellenecek alanı belirleyerek değer atama
	err = services.UpdateRecordField(record, field, fileKey)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Veritabanı güncelleme hatası")
	}

	// Yanıt döndür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":  "Dosya başarıyla yüklendi ve veritabanına kaydedildi",
		"file_key": fileKey,
		"updated":  record,
	})
}
