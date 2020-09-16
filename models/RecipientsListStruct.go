package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type RecipientsList struct {
	ID         uint64       `gorm: "primary_key"`
	Name       string       `gorm:"size:255"`
	Recipients []Recipient
	OrganismId uint64
}

func FindRecipientsListByID(uid uint32) (RecipientsList, error) {
	var err error
	var recipientList RecipientsList
	err = db.Debug().Model(&RecipientsList{}).Where("id = ?", uid).Take(&recipientList).Error
	if err != nil {
		return RecipientsList{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return RecipientsList{}, errors.New("Recipients List Not Found")
	}
	return recipientList, nil
}

func FindRecipientsListsByOrganismID(id uint64) ([]RecipientsList, error) {
	var err error
	var recipientslist []RecipientsList
	err = db.Debug().Where("organism_id = ?", id).Find(&recipientslist).Error
	if err != nil {
		return nil, err
	}
	return recipientslist, nil
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

func DeleteRecipientsListByID(id uint64) error {
	var err error
	var recipientList RecipientsList

	err = db.Debug().First(&recipientList, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Recipient List Not Found")
	}
	err = db.Debug().Delete(&recipientList, id).Error
	if err != nil {
		return err
	}

	return nil
}

func EditRecipientsListByID(recipentList *RecipientsList, id uint64) error {
	var err error
	var old RecipientsList
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Recipient List Not Found")
	}
	recipentList.ID = id
	err = db.Debug().Save(&recipentList).Error
	if err != nil {
		return errors.New("Could'nt update recipient list")
	}
	return nil
}

func CreateRecipientList(recipientList *RecipientsList) error {
	var err error
	err = db.Debug().Model(RecipientsList{}).Create(recipientList).Error

	if err != nil {
		return err
	}

	return nil
}
