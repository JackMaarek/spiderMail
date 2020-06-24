package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Registration(c *gin.Context){
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var err = models.ValidateUser(user,"")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Println(user)
	userCreated, err := models.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, "User has been created:" + userCreated.Email)
}

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := models.ValidateUser(u,"login")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide valid login details")
		return
	}

	token, err := services.SignIn(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide valid credentials")
		return
	}

	c.JSON(http.StatusOK, token)
}