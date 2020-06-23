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

	all_users := db.Find(&users)
	if all_users.Error != nil {
		panic(all_users.Error)
	}

	return all_users
}

func GetUserById(id int) *gorm.DB{
	db := Database.Connect()
	defer db.Close()

	// Get user by id
	var user User

	res_user := db.First(&user, id)
	if res_user.Error != nil {
		panic(res_user.Error)
	}

	return res_user
}

func CreateUser(user User) *gorm.DB{
	db := Database.Connect()
	defer db.Close()

	res_user := db.Create(user)
	if res_user.Error != nil {
		panic(res_user.Error)
	}

	return res_user
}

func DeleteUserById (id int) *gorm.DB{
	db := Database.Connect()
	defer db.Close()
	var user User

	response := db.Delete(&user, id)
	if response.Error != nil {
		panic(response.Error)
	}

	return response
}