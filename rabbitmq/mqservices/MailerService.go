package mqservices

import (
	"github.com/caarlos0/env/v6"
	"gopkg.in/gomail.v2"
	"log"
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
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println(err)
		return err
	}

	err := sendMail("902468820252-ta6b9t3tcnon0mfm4b7cvcj6aokgbn5r.apps.googleusercontent.com", "GPpl0cxKVS75dMuIo_mP2YP4", mail.Recipient, mail.Subject, mail.Body)
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

	d := gomail.NewDialer("smtp.gmail.com", 465, author, password)

	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}
