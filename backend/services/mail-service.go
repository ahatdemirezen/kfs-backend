package services

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/mailjet/mailjet-apiv3-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateVerificationCode() string {
	code := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", code)
}

func SendVerificationEmail(recipientEmail, verificationCode string) error {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MAIL_JET_API_KEY"), os.Getenv("MAIL_JET_SECRET_KEY"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "your-verified-sender@domain.com",
				Name:  "KFS Platform",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: recipientEmail,
				},
			},
			Subject:  "Email Doğrulama Kodu",
			TextPart: "Doğrulama Kodunuz: " + verificationCode,
			HTMLPart: fmt.Sprintf("<h3>Email Doğrulama</h3><br />Doğrulama kodunuz: <strong>%s</strong>", verificationCode),
		},
	}

	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	return err
} 