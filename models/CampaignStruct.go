package models

import (
	"github.com/jinzhu/gorm"
	"time"
	"errors"
)

type Campaign struct {
	ID          uint64    `gorm:"primary_key"`
	Name        string    `gorm:"size:255"`
	DateCreated time.Time `gorm:"default:NULL ON UPDATE CURRENT_TIMESTAMP"`
	OrganismId uint64
	Subject     string    `gorm:"size:255"`
	Content 	string 	  `gorm:"size:1023"`
	RecipientsListId  uint64
}

func FindCampaignsByCompany(uid uint32) (*[]Campaign, error) {
	var err error
	var campaigns []Campaign

	err = db.Debug().Model(Campaign{}).Where("company_id = ?", uid).Take(&campaigns).Error

	if err != nil {
		return &[]Campaign{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Campaign{}, errors.New("No campaigns found")
	}
	return &campaigns, err
}

func FindCampaignByID(uid uint32) (*Campaign, error) {
	var err error
	var campaign Campaign
	err = db.Debug().Model(Campaign{}).Where("id = ?", uid).Take(&campaign).Error
	if err != nil {
		return &Campaign{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Campaign{}, errors.New("Campaign Not Found")
	}
	return &campaign, err
}

func DeleteCampaignByID(uid uint32) (*Campaign, error) {
	var err error
	var campaign Campaign

	err = db.Debug().Model(Campaign{}).Where("id = ?", uid).Delete(&campaign).Error
	if err != nil {
		return &Campaign{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Campaign{}, errors.New("Campaign Not Found")
	}
	return &campaign, err
}

func EditCampaignByID(campaign Campaign) (*Campaign, error) {
	var err error
	uid := campaign.ID

	err = db.Debug().Model(Campaign{}).Where("id = ?", uid).Save(&campaign).Take(&campaign).Error
	if err != nil {
		return &Campaign{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Campaign{}, errors.New("Campaign Not Found")
	}
	return &campaign, err
}

func CreateCampaign(campaign Campaign) (*Campaign, error) {
	var err error
	err = db.Debug().Model(Campaign{}).Create(campaign).Error

	if err != nil {
		return &Campaign{}, err
	}

	return &campaign, err
}