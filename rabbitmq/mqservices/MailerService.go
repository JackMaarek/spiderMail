package mqservices

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
	"gopkg.in/gomail.v2"
)

type Config struct {
	PROVIDER_KEY    string `env:"PROVIDER_KEY"`
	PROVIDER_SECRET string `env:"PROVIDER_SECRET"`
}

type Mail struct {
	Recipient string
	Subject   string
	Body      string
}

func CallMailerService(mail *Mail) error {
	fmt.Println("sdfgfsgh")
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println(err)
		return err
	}

	var password string = "contact.jason.gauvin@gmail.com"
	var author string = "uifbpxvvqwttehyy"

	err := sendMail(author, password, "contact.jason.gauvin@gmail.com", mail.Subject, mail.Body)
	if err != nil {
		return err
	}
	return nil
}

func sendMail(author string, password string, to string, subject string, body string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", author)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)
	fmt.Println(author)
	fmt.Println(password)

	d := gomail.NewDialer("smtp.gmail.com", 465, author, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
