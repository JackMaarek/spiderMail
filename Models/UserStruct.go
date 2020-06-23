package Models

import (
	"github.com/JackMaarek/spiderMail/Database"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID uint64         `gorm:"primary_key"`
	Name string       `gorm:"size:255"`
	Password string   `gorm:"size:255"`
	Email string      `gorm:"size:255"`
	//Company *Organism //`json:"company",one-to-many:"Organism"`
}

func GetAllUsers() *gorm.DB{
	db := Database.Connect()
	defer db.Close()

	// Get users
	var users []User

	return db.Find(&users)
}

func GetUserById(id int) *gorm.DB{
	db := Database.Connect()
	defer db.Close()

	// Get user by id
	var user User

	return db.First(&user, 1)
}

func CreateUser(user User) *gorm.DB{
	db := Database.Connect()
	defer db.Close()

}