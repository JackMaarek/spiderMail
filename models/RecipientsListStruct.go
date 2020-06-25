package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type RecipientsList struct {
	ID         uint64       `gorm: "primary_key"`
	Name       string       `gorm:"size:255"`
	Recipients []*Recipient `gorm:"many2many:mail_lists;"`
	OrganismId uint64
}

func FindRecipientsListByID(uid uint32) (*RecipientsList, error) {
	var err error
	var recipientList *RecipientsList
	err = db.Debug().Model(RecipientsList{}).Where("id = ?", uid).Take(&recipientList).Error
	if err != nil {
		return nil, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &RecipientsList{}, errors.New("Recipients List Not Found")
	}
	return nil, err
}

func FindRecipientsList() ([]RecipientsList, error) {
	var err error
	var recipientsList []RecipientsList
	err = db.Debug().Find(&recipientsList).Error
	if err != nil {
		return nil, err
	}
	return recipientsList, nil
}

func DeleteRecipientsListByID(uid uint32) (*RecipientsList, error) {
	var err error
	var recipientList RecipientsList

	err = db.Debug().Model(RecipientsList{}).Where("id = ?", uid).Delete(&recipientList).Error
	if err != nil {
		return &RecipientsList{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &RecipientsList{}, errors.New("Recipients List Not Found")
	}
	return &recipientList, err
}

func EditRecipientsListByID(recipientList RecipientsList) (*RecipientsList, error) {
	var err error
	uid := recipientList.ID

	err = db.Debug().Model(RecipientsList{}).Where("id = ?", uid).Save(&recipientList).Take(&recipientList).Error
	if err != nil {
		return &RecipientsList{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &RecipientsList{}, errors.New("Recipients List Not Found")
	}
	return &recipientList, err
}

func CreateRecipientList(recipientList *RecipientsList) error {
	var err error
	err = db.Debug().Model(RecipientsList{}).Create(recipientList).Error

	if err != nil {
		return err
	}

	return nil
}
