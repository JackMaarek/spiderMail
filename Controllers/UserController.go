package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/JackMaarek/spiderMail/Models"
	"github.com/jinzhu/gorm"
	"net/http"
)


func GetUsers(c *gin.Context) {
	db, err := gorm.Open("mysql", "spidermail:spidermail@(localhost:3306)/spidermail?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&Models.User{})
	var user Models.User
	users := db.Limit(3).Find(&user)

	c.JSON(http.StatusOK, users)
}