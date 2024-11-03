package main

import (
	"crypto/tls"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/mail.v2"
)

var (
	MailHost = os.Getenv("MAIL_HOST")
	MailFrom = os.Getenv("MAIL_FROM")
	MailPort = 587
	MailUser = os.Getenv("MAIL_USER")
	MailPass = os.Getenv("MAIL_PASS")
)

func main() {
	message := mail.NewMessage()

	message.SetHeader("From", MailFrom)
	message.SetHeader("To", "example@mail.com") // email sent to
	message.SetHeader("Subject", "gomail")

	// message.SetBody("text/plain", "Hello from gomail")
	message.SetBody("text/html", `
    <html>
      <body>
        <h1>This is a Test Email</h1>
        <p><b>Hello!</b> Please find the attachment below.</p>
        <p>Thanks,gomail</p>
      </body>
    </html>
  `)
	// message.Attach("public/invoice.pdf")

	dialer := mail.NewDialer(MailHost, MailPort, MailUser, MailPass)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
