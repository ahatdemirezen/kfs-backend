package middleware

import (
	"log"
	"github.com/gofiber/fiber/v2"
)
//superAdminAuthMiddleware, kullanıcının superadmin olup olmadığını kontrolü yapar

func SuperAdminAuthMiddleware(c *fiber.Ctx) error {
	roles , ok:= c.Locals("roles").([]string)
	if !ok {
		log.Println("HATA: Rol bilgisi bulunamadı veya yanlış formatta")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Rol bilgisi bulunamadı",
		})
	}

	//kullanıcının "superadmin" rolü  olup olmadığını kontrol et

	isSuperAdmin := false
	for _,role := range roles {
		if role=="superadmin" {
			isSuperAdmin = true
			break
		}
}
	if !isSuperAdmin {
		log.Println("HATA: yetkisiz erişim - Kullanıcı superadmin değil")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error" : "BU endpoint için Yetkiniz bulunmamaktadır",
	})
}
log.Println("INFO: Superadmin yetkilendirildi - Roller: %v", roles)
return c.Next()
}
