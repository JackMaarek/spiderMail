package Models

import (
	"os"
	"time"
)

type Campaign struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	DateCreated time.Time `json:"date_created"`
	User        User      //`json:"user"`
	Subject     string    `json:"subject"`
	Attachments os.File   `json:"attachments"`
	Recipients  []Recipient
}
