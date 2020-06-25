package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Campaign struct {
	ID               uint64    `gorm:"primary_key"`
	Name             string    `gorm:"size:255"`
	DateCreated      time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	OrganismId       uint64
	Subject          string `gorm:"size:255"`
	Content          string `gorm:"size:1023"`
	RecipientsListId uint64
}

func FindCampaigns() ([]Campaign, error) {
	var err error
	var campaigns []Campaign
	err = db.Debug().Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func FindCampaignsByOrganismID(id uint64) ([]Campaign, error) {
	var err error
	var campaigns []Campaign
	err = db.Debug().Where("organism_id = ?", id).Find(&campaigns).Error
	if err != nil {
		return nil, err
	}
	return campaigns, nil
}

func FindCampaignByID(uid uint64) (Campaign, error) {
	var err error
	var campaign Campaign
	err = db.Debug().First(&campaign, uid).Error
	if err != nil {
		return Campaign{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Campaign{}, errors.New("Organism Not Found")
	}
	return campaign, nil
}

func DeleteCampaignByID(id uint64) error {
	var err error
	var campaign Campaign

	err = db.Debug().First(&campaign, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Campaign Not Found")
	}
	err = db.Debug().Delete(&campaign, id).Error
	if err != nil {
		return err
	}

	return nil
}

func EditCampaignByID(campaign *Campaign, id uint64) error {
	var err error
	var old Campaign
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Organism Not Found")
	}
	campaign.ID = id
	campaign.DateCreated = old.DateCreated
	err = db.Debug().Save(&campaign).Error
	if err != nil {
		return errors.New("Could'nt update organism")
	}
	return nil
}

func CreateCampaign(campaign *Campaign) error {
	var err error
	campaign.DateCreated = time.Now()
	err = db.Debug().Create(campaign).Error

	if err != nil {
		return err
	}
	return nil
}
