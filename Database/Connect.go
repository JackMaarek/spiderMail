package Database

import "github.com/jinzhu/gorm"

func Connect() *gorm.DB{
	db, err := gorm.Open("mysql", "spidermail:spidermail@tcp(mysql_db)/spidermail?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	return db
}