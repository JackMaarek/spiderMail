package controllers

import (
	"fmt"
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

func DeleteUser(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error

	_, err = models.DeleteUserByID(id)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusOK, "User has been deleted")
}
