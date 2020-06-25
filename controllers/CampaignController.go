package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	var err error
	campaigns, err = models.FindCampaigns()

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusOK, campaigns)
}

func GetCampaignById(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error
	var campaign models.Campaign

	campaign, err = models.FindCampaignByID(id)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusOK, campaign)
}

func GetCampaignsByOrganismId(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error
	var campaigns []models.Campaign

	campaigns, err = models.FindCampaignsByOrganismID(id)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusOK, campaigns)
}
func DeleteCampaignById(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var err error
	err = models.DeleteCampaignByID(id)

	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return
	}

	message := "Campaign with id " + c.Param("id") + " have been deleted."
	c.JSON(http.StatusOK, message)
}

func EditCampaignById(c *gin.Context) {
	// Get id and converts it
	id := services.ConvertStringToInt(c.Param("id"))

	var campaign models.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var err error
	err = models.EditCampaignByID(&campaign, id)

	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return
	}

	c.JSON(http.StatusOK, campaign)
}

func CreateCampaign(c *gin.Context) {
	var campaign models.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := models.CreateCampaign(&campaign)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusCreated, campaign)
}
