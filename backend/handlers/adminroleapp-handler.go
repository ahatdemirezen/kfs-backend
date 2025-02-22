package handlers

import (
	"kfs-backend/services"
	"strconv"
	"github.com/gofiber/fiber/v2"
)

var roleService = services.NewRoleApplicationService()

func AdminRoleApplicationsHandler(c *fiber.Ctx) error {
	roles, ok := c.Locals("roles").([]string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Rol bilgisi bulunamadı")
	}

	isAuthorized := false
	for _, role := range roles {
		if role == "admin" || role == "superadmin" {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		return c.Status(fiber.StatusForbidden).SendString("Yetkisiz erişim")
	}

	applications, err := roleService.GetAllRoleApplications()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Başvurular getirilemedi")
	}
	return c.JSON(applications)
}

func UpdateRoleApplicationStatusHandler(c *fiber.Ctx) error {
	roles, ok := c.Locals("roles").([]string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Rol bilgisi bulunamadı")
	}

	isAuthorized := false
	for _, role := range roles {
		if role == "admin" || role == "superadmin" {
			isAuthorized = true
			break
		}
	}

	if !isAuthorized {
		return c.Status(fiber.StatusForbidden).SendString("Yetkisiz erişim")
	}

	applicationIdStr := c.Params("id")
	applicationId, err := strconv.ParseUint(applicationIdStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz başvuru ID")
	}

	var requestBody struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz istek verisi")
	}

	if requestBody.Status != "accepted" && requestBody.Status != "rejected" && requestBody.Status != "pending" {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz status değeri")
	}

	err = roleService.UpdateRoleApplicationStatus(uint(applicationId), requestBody.Status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Başvuru statüsü güncellenemedi")
	}
	return c.Status(fiber.StatusOK).SendString("Başvuru statüsü başarıyla güncellendi")
}
