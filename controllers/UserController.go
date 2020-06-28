package controllers

import (
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateUser(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	user, findError := models.FindUserByID(id)
	if findError != nil {
		c.JSON(http.StatusUnprocessableEntity, findError.Error())
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	userUpdated, err := models.EditUserByID(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "User has been updated: "+userUpdated.Name+userUpdated.Email)
}

func GetUsersByOrganism(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error
	var users []models.User

	users, err = models.FindUsersByOrganismID(id)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Unable to update user")
		return
	}

	c.JSON(http.StatusOK, users)
}

func DeleteUser(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error

	_, err = models.DeleteUserByID(id)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Unable to delete user")
		return
	}

	c.JSON(http.StatusOK, "User has been deleted")
}

func RefreshToken(c *gin.Context)  {
	var token string
	var err error
	token = services.ExtractToken(c)
	err = models.UpdateToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, "Token has been refreshed")
}