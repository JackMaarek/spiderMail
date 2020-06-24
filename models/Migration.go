package models

func MakeMigrations() {
	db.AutoMigrate(&User{}, &Campaign{}, &Organism{}, &Recipient{}, &Token{})
}
