package models

import (
	"github.com/jinzhu/gorm"

	"errors"
)

type Organism struct {
	ID uint64   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}


func FindOrganismByID(uid uint32) (*Organism, error) {
	var err error
	var organism Organism
	err = db.Debug().Model(Organism{}).Where("id = ?", uid).Take(&organism).Error
	if err != nil {
		return &Organism{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Organism{}, errors.New("Organism Not Found")
	}
	return &organism, err
}

func DeleteOrganismByID(uid uint32) (*Organism, error) {
	var err error
	var organism Organism

	err = db.Debug().Model(Organism{}).Where("id = ?", uid).Delete(&organism).Error
	if err != nil {
		return &Organism{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Organism{}, errors.New("Organism Not Found")
	}
	return &organism, err
}

func EditOrganismByID(organism Organism) (*Organism, error) {
	var err error
	uid := organism.ID

	err = db.Debug().Model(Campaign{}).Where("id = ?", uid).Save(&organism).Take(&organism).Error
	if err != nil {
		return &Organism{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Organism{}, errors.New("Organism Not Found")
	}
	return &organism, err
}

func CreateOrganism(organism Organism) (*Organism, error) {
	var err error
	err = db.Debug().Model(Organism{}).Create(organism).Error

	if err != nil {
		return &Organism{}, err
	}

	return &organism, err
}