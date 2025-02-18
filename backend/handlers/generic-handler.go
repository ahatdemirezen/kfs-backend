package handlers

import (
	
	"kfs-backend/services"
	"kfs-backend/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateVenture[T any, R any](c *fiber.Ctx, service services.VentureService[T], req R) error {
	// JSON verisini parse et
	if err := c.BodyParser(&req); err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz JSON formatı")
	}

	// Burada doğru tür parametrelerini belirtin
	model := utils.ConvertRequestToModel[R, T](req)

	// Service'i kullanarak kaydı oluştur
	createdRecord, err := service.Create(&model)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kayıt oluşturulurken hata oluştu")
	}

	// Başarılı yanıt döndür
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Kayıt başarıyla oluşturuldu",
		"data":    createdRecord,
	})
}


func GetVentureByID[T any](c *fiber.Ctx, service services.VentureService[T]) error {
	// ID parametresini al
	id, err := utils.GetIDParam(c, "id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Service'i kullanarak kaydı getir
	record, err := service.GetByID(id)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusNotFound, "Kayıt bulunamadı")
	}

	// Başarılı yanıt döndür
	return c.Status(fiber.StatusOK).JSON(record)
}


func GetVenturesByField[T any](c *fiber.Ctx, service services.VentureService[T], field string) error {
    // Parametre değerini al
    value, err := utils.GetIDParam(c, field) // Örneğin, "product_model_info_id" parametresini al
    if err != nil {
        return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz parametre değeri")
    }

    // Service'i kullanarak kayıtları getir
    records, err := service.GetByField(field, value)
    if err != nil {
        return utils.RespondWithError(c, fiber.StatusNotFound, "Kayıtlar bulunamadı")
    }

    // Başarılı yanıt döndür
    return c.Status(fiber.StatusOK).JSON(records)
}


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

	// Request yapısını model yapısına dönüştür
	model := utils.ConvertRequestToModel[R, T](req)


	// Service'i kullanarak kaydı güncelle
	updatedRecord, err := service.Update(id, &model)
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kayıt güncellenirken hata oluştu")
	}

	// Başarılı yanıt döndür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Kayıt başarıyla güncellendi",
		"data":    updatedRecord,
	})
}

func DeleteVenture[T any](c *fiber.Ctx, service services.VentureService[T]) error {
	// ID parametresini al
	id, err := utils.GetIDParam(c, 	"id")
	if err != nil {
		return utils.RespondWithError(c, fiber.StatusBadRequest, "Geçersiz ID")
	}

	// Service'i kullanarak kaydı sil
	if err := service.Delete(id); err != nil {
		return utils.RespondWithError(c, fiber.StatusInternalServerError, "Kayıt silinirken hata oluştu")
	}

	// Başarılı yanıt döndür
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Kayıt başarıyla silindi"})
}

