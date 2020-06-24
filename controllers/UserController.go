package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateUser(id string, c *gin.Context){
	userId , converr := strconv.ParseUint(id, 10, 32)
	if converr != nil {
		c.JSON(http.StatusUnprocessableEntity, converr.Error())
		return
	}

	user, findError := models.FindUserByID(userId)
	if findError != nil {
		c.JSON(http.StatusUnprocessableEntity, findError.Error())
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var err = models.ValidateUser(*user,"update")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	fmt.Println(user)
	userUpdated, err := models.EditUser(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "User has been updated: " + userUpdated.Name + userUpdated.Email)
}

func DeleteUser(id string, c *gin.Context) {
	userId, converr := strconv.ParseUint(id, 10, 32)
	if converr != nil {
		c.JSON(http.StatusUnprocessableEntity, converr.Error())
		return
	}
	user, findError := models.FindUserByID(userId)
	if findError != nil {
		c.JSON(http.StatusUnprocessableEntity, findError.Error())
		return
	}
	_, deleteErr := models.DeleteSingleUser(user)
	if deleteErr != nil {
		c.JSON(http.StatusUnprocessableEntity, deleteErr.Error())
		return
	}

	c.JSON(http.StatusOK, "User has been deleted")
}