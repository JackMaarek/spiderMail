package Models

import "github.com/JackMaarek/spiderMail/Database"

func MakeMigrations() {
	db := Database.Connect()

	// Make migrations
	db.AutoMigrate(&User{}, &Campaign{}, &Organism{}, &Recipient{}, &Token{})

	db.Close()
}