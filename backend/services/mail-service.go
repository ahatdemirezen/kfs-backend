package services

import (
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendVerificationEmail(recipientEmail string, verificationCode string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "SendGrid API anahtarı bulunamadı")
	}

	fromEmail := "ahat@jxpuniworkhub.com"
	from := mail.NewEmail("", fromEmail)
	subject := "Email Doğrulama Kodu"
	to := mail.NewEmail("", recipientEmail)

	plainTextContent := "Doğrulama kodunuz: " + verificationCode

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(apiKey)

	response, err := client.Send(message)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Email gönderme hatası: "+err.Error())
	}

	if response.StatusCode >= 400 {
		return fiber.NewError(fiber.StatusInternalServerError, "Email gönderme başarısız. Status: "+strconv.Itoa(response.StatusCode))
	}

	return nil
}
