package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/midlewares"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Registration(c *gin.Context){
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var err = models.ValidateUser(&user,"")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	userCreated, err := models.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	var tokenCreated *models.Token
	tokenCreated, err = models.CreateTokenFromUser(userCreated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(tokenCreated)
	c.JSON(http.StatusCreated, "User has been created:" + userCreated.Name + userCreated.Email)
}

func Login(c *gin.Context) {
	if err := midlewares.CheckAuthorization(c); err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := models.ValidateUser(&u,"login")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide valid login details")
		return
	}

	err = SignIn(u.Email, u.Password)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide valid credentials")
		return
	}

	c.JSON(http.StatusOK, "Successfully signed in.")
}

func SignIn(email string, password string)  error {

	var err error

	var user *models.User

	user, err = models.FindUserByEmail(email)
	if err != nil {
		return err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	return nil
}