package models

import (
	"github.com/jinzhu/gorm"
)

type Recipient struct {
	Name  string `gorm:"size:255"`
	Email string `gorm:"size:255"`
}

func GetRecipients() *gorm.DB {

	var recipients Recipient
	all_recipients := db.Find(&recipients)

	if all_recipients.Error != nil {
		panic(all_recipients.Error)
	}

	return all_recipients
}

func GetRecipientById(id int) *gorm.DB {

	var recipient Recipient
	res_recipient := db.First(&recipient, id)
	if res_recipient.Error != nil {
		panic(res_recipient.Error)
	}

	return res_recipient
}

func CreateRecipient(recipient Recipient) *gorm.DB {

	res_recipient := db.Create(recipient)
	if res_recipient.Error != nil {
		panic(res_recipient.Error)
	}

	return res_recipient
}

func DeleteRecipientbyId(id int) *gorm.DB {

	var recipient Recipient
	response := db.Delete(&recipient, id)
	if response.Error != nil {
		panic(response.Error)
	}

	return response
}