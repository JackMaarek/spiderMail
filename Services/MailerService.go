package Services

import (
	"github.com/caarlos0/env/v6"
	"gopkg.in/gomail.v2"
	"os"
)

type Config struct {
	GMAIL_USER   	string `env:"GMAIL_USER"`
	GMAIL_PASSWORD  string    `env:"GMAIL_PASSWORD"`
}

func main() {
	cfg := Config{}
	env.Parse(&cfg)

	var password string = os.Getenv("GMAIL_PASSWORD")
	var author string = os.Getenv("GMAIL_USER")

	sendMail(author, password, "edwin.vautier@gmail.com", "Test", "<h1>Hello Edwin</h1><br><ul><li>1</li><li>2</li></ul>")
}

func sendMail(author string,password string, to string, subject string, body string) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", author)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	d := gomail.NewDialer("smtp.gmail.com", 465, author, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(mail); err != nil {
		panic(err)
	}
}
