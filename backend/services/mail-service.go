package services

import (
	"fmt"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendVerificationEmail(recipientEmail string, verificationCode string) error {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("SendGrid API anahtarı bulunamadı")
	}

	fmt.Printf("SendGrid API Key: %s\n", apiKey[:10]) // API key'in sadece ilk 10 karakterini göster

	fromEmail := "ahat@jxpuniworkhub.com"
	
	fmt.Printf("Gönderen email: %s\n", fromEmail)
	
	from := mail.NewEmail("", fromEmail) // İsim boş string olarak verildi
	subject := "Email Doğrulama Kodu"
	to := mail.NewEmail("", recipientEmail) // Alıcı ismi de boş string
	
	plainTextContent := fmt.Sprintf("Doğrulama kodunuz: %s", verificationCode)
	
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, "")
	client := sendgrid.NewSendClient(apiKey)
	
	fmt.Printf("Email gönderme isteği hazırlandı. Alıcı: %s\n", recipientEmail)
	
	response, err := client.Send(message)
	if err != nil {
		fmt.Printf("SendGrid Hata: %v\n", err)
		return fmt.Errorf("email gönderme hatası: %v", err)
	}
	
	fmt.Printf("SendGrid Response - Status: %d, Body: %s\n", response.StatusCode, response.Body)
	
	if response.StatusCode >= 400 {
		return fmt.Errorf("email gönderme başarısız. Status: %d, Body: %s", response.StatusCode, response.Body)
	}
	
	return nil
}