package models

import (
	"github.com/jinzhu/gorm"

	"errors"
)

type Recipient struct {
	ID               uint64 `gorm: "primary_key"`
	Name             string `gorm:"size:255"`
	Email            string `gorm:"size:255"`
	RecipientsList   RecipientsList
	RecipientsListID uint64
}

func FindRecipientsByListId(uid uint32) (*[]Recipient, error) {
	var err error
	var recipients []Recipient

	var (
		ID uint64
		Name string
		Email string
	)
	rows, _ := db.Raw("SELECT * FROM recipients WHERE recipients_list_id = ?", uid).Rows()
	for rows.Next() {
		rows.Scan(&ID, &Name, &Email)
		recipient := Recipient {
			ID: ID,
			Name: Name,
			Email: Email,
		}
		recipients = append(recipients, recipient)
	}
	if err != nil {
		return &[]Recipient{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Recipient{}, errors.New("No recipients found")
	}
	return &recipients, err
}

func FindRecipientByID(uid uint32) (*Recipient, error) {
	var err error
	var recipient Recipient
	err = db.Debug().Model(Recipient{}).Where("id = ?", uid).Take(&recipient).Error
	if err != nil {
		return &Recipient{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Recipient{}, errors.New("Recipient Not Found")
	}
	return &recipient, err
}

func DeleteRecipientByID(uid uint32) (*Recipient, error) {
	var err error
	var recipient Recipient

	err = db.Debug().Model(Campaign{}).Where("id = ?", uid).Delete(&recipient).Error
	if err != nil {
		return &Recipient{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Recipient{}, errors.New("Recipient Not Found")
	}
	return &recipient, err
}

func EditrecipientByID(recipient Recipient) (*Recipient, error) {
	var err error
	uid := recipient.ID

	err = db.Debug().Model(Recipient{}).Where("id = ?", uid).Save(&recipient).Take(&recipient).Error
	if err != nil {
		return &Recipient{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Recipient{}, errors.New("Recipient Not Found")
	}
	return &recipient, err
}

func CreateRecipient(recipient *Recipient) (*Recipient, error) {
	var err error
	err = db.Debug().Model(Recipient{}).Create(recipient).Error

	if err != nil {
		return &Recipient{}, err
	}

	return recipient, err
}
