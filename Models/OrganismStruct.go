package Models

import (
	"github.com/JackMaarek/spiderMail/Database"
	"github.com/jinzhu/gorm"
)

type Organism struct {
	ID uint64   `gorm:"primary_key"`
	Name string `gorm:"size:255"`
}

func GetOrganisms() *gorm.DB {
	db := Database.Connect()
	defer db.Close()

	var organisms []Organism
	all_organisms := db.Find(&organisms)
	if all_organisms.Error != nil {
		panic(all_organisms.Error)
	}

	return all_organisms
}

func GetOrganismById(id int) *gorm.DB {
	db := Database.Connect()
	defer db.Close()

	var organism Organism
	res_organism := db.First(&organism, id)
	if res_organism.Error != nil {
		panic(res_organism.Error)
	}

	return res_organism
}

func CreateOrganism(organism Organism) *gorm.DB {
	db := Database.Connect()
	defer db.Close()

	res_organism := db.Create(organism)
	if res_organism.Error != nil {
		panic(res_organism.Error)
	}

	return res_organism
}

func DeleteOrganismbyId(id int) *gorm.DB {
	db := Database.Connect()
	defer db.Close()

	var organism Organism
	response := db.Delete(&organism, id)
	if response.Error != nil {
		panic(response.Error)
	}

	return response
}