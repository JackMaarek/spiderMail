package models

import (
	"time"
)

type Campaign struct {
	ID          uint64    `gorm:"primary_key"`
	Name        string    `gorm:"size:255"`
	DateCreated time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	User        User
	Subject     string    `gorm:"size:255"`
	Content 	string 	  `gorm:"size:1023"`
	//Attachments os.File   `gorm:"attachments"`
	Recipients  []Recipient
}
