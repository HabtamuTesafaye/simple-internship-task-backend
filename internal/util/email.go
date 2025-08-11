package util

import (
	"os"

	"gopkg.in/mail.v2"
)

func SendEmail(to, subject, htmlBody string) error {
	message := mail.NewMessage()
	message.SetHeader("From", os.Getenv("EMAIL_FROM"))
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", htmlBody)

	dialer := mail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		587,
		os.Getenv("EMAIL_USER"),
		os.Getenv("EMAIL_PASS"),
	)

	return dialer.DialAndSend(message)
}
