package service

import (
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

func MailService() *mail.SMTPClient {

	server := mail.NewSMTPClient()
	server.Host = "smtp.mailtrap.io"
	server.Port = 587
	server.Username = "7d3316c52e22c1"
	server.Password = "f0179338ccd69a"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()

	if err != nil {
		log.Fatal(err)
	}

	return smtpClient
}
