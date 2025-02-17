package handlers

import (
	"kfs-backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

func AdminRoleApplicationsHandler(c *fiber.Ctx) error {
	role := c.Get("Role")
	if role != "admin" && role != "superadmin" {
		return c.Status(fiber.StatusForbidden).SendString("Yetkisiz erişim")
	}
	applications, err := services.GetAllRoleApplications()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Başvurular getirilemedi")
	}
	return c.JSON(applications)
}

func UpdateRoleApplicationStatusHandler(c *fiber.Ctx) error {
	role := c.Get("Role")
	if role != "admin" && role != "superadmin" {
		return c.Status(fiber.StatusForbidden).SendString("Yetkisiz erişim")
	}
	applicationIdStr := c.Params("id")
	applicationId, err := strconv.Atoi(applicationIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz başvuru ID")
	}
	var requestBody struct { Status string `json:"status"` }
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz istek verisi")
	}
	if requestBody.Status != "accepted" && requestBody.Status != "rejected" && requestBody.Status != "pending" {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz status değeri")
	}
	err = services.UpdateRoleApplicationStatus(applicationId, requestBody.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Başvuru statüsü güncellenemedi")
	}
	return c.Status(fiber.StatusOK).SendString("Başvuru statüsü başarıyla güncellendi")
}
