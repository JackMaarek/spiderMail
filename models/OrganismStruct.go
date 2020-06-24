package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type Organism struct {
	ID   uint64 `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

func FindOrganismByID(uid uint64) (Organism, error) {
	var err error
	var organism Organism
	err = db.Debug().First(&organism, uid).Error
	if err != nil {
		return Organism{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return Organism{}, errors.New("Organism Not Found")
	}
	return organism, nil
}

func FindOrganisms() ([]Organism, error) {
	var err error
	var organisms []Organism
	err = db.Debug().Find(&organisms).Error
	if err != nil {
		return nil, err
	}
	return organisms, nil
}

func DeleteOrganismByID(id uint64) error {
	var err error
	var organism Organism

	err = db.Debug().First(&organism, id).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Organism Not Found")
	}
	err = db.Debug().Delete(&organism, id).Error
	if err != nil {
		return err
	}

	return nil
}

func EditOrganismByID(organism *Organism, id uint64) error {
	var err error
	var old Organism
	err = db.Debug().Where("id = ?", id).First(&old).Error
	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Organism Not Found")
	}
	organism.ID = id
	err = db.Debug().Save(&organism).Error
	if err != nil {
		return errors.New("Could'nt update organism")
	}
	return nil
}

func CreateOrganism(organism *Organism) error {
	var err error
	err = db.Debug().Create(organism).Error

	if err != nil {
		return err
	}
	return nil
}
