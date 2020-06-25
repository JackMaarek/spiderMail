package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRecipientList(c *gin.Context) {
	var recipientsList []models.RecipientsList
	var err error
	recipientsList, err = models.FindRecipientsList()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusOK, recipientsList)
}

func CreateRecipientsList(c *gin.Context) {
	var recipientsList models.RecipientsList
	if err := c.ShouldBindJSON(&recipientsList); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := models.CreateRecipientList(&recipientsList)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusCreated, recipientsList)

}
