package main

import (
	"github.com/caarlos0/env/v6"
	"gopkg.in/gomail.v2"
	"os"
)

type Config struct {
	PROVIDER_KEY   	string `env:"PROVIDER_KEY"`
	PROVIDER_SECRET  string    `env:"PROVIDER_SECRET"`
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	var password string = os.Getenv("PROVIDER_SECRET")
	var author string = os.Getenv("PROVIDER_KEY")

	sendMail(author, password, "edwin.vautier@gmail.com", "Test", "<h1>Hello Edwin</h1><br><ul><li>1</li><li>2</li></ul>")
}

func sendMail(author string,password string, to string, subject string, body string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", author)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 465, author, password)

	if err := d.DialAndSend(mail); err != nil {
		panic(err)
	}
}
