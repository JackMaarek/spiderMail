package controllers

import (
	"github.com/JackMaarek/spiderMail/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Registration(c *gin.Context) {
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	var err = models.ValidateUser(&user, "Please provide valid login details")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	userCreated, err := models.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, "An error occurred")
		return
	}
	// @fixme: Should send email to confirm user registration.
	var tokenCreated *models.Token
	tokenCreated, err = models.CreateTokenFromUser(userCreated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "An error occurred")
		return
	}

	c.Header("Authorization", "Bearer: "+tokenCreated.Token)
	c.JSON(http.StatusCreated, "User has been created:"+userCreated.Name+userCreated.Email)
}

func Login(c *gin.Context) {
	var u models.User
	var tokenString string
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	err := models.ValidateUser(&u, "login")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide valid login details")
		return
	}

	tokenString, err = SignIn(&u)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Please provide valid credentials")
		return
	}
	c.Header("Authorization", "Bearer: " + tokenString)
	c.JSON(http.StatusOK, u)
}

func SignIn(u *models.User) (string, error) {

	var err error
	var lu *models.User
	lu, err = models.FindUserByEmail(u.Email)
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(lu.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	var token *models.Token
	token, err = models.FindTokenByUserID(lu.ID)
	if err != nil {
		return "", err
	}
	
	u.ID = lu.ID
	u.OrganismId = lu.OrganismId

	return token.Token, nil
}
