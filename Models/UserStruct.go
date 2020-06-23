package Models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID uint64         `gorm:"primary_key"`
	Name string       `gorm:"size:255"`
	Password string   `gorm:"size:255"`
	Email string      `gorm:"size:255"`
	//Company *Organism //`json:"company",one-to-many:"Organism"`
}
