package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/services"
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

func GetRecipientsListById(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error
	var recipientsList models.RecipientsList

	recipientsList, err = models.FindRecipientsListByID(uint32(id))

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusOK, recipientsList)
}

func EditRecipientsListById(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var recipientList models.RecipientsList
	if err := c.ShouldBindJSON(&recipientList); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var err error
	err = models.EditRecipientsListByID(&recipientList, id)

	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return
	}

	c.JSON(http.StatusOK, recipientList)
}

func DeleteRecipientsListById(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error
	err = models.DeleteRecipientsListByID(id)

	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return
	}

	message := "Recipient List with id " + c.Param("id") + " have been deleted."
	c.JSON(http.StatusOK, message)
}
