package controllers

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrganisms(c *gin.Context) {
	var organisms []models.Organism
	var err error
	organisms, err = models.FindOrganisms()

	if err != nil {
		fmt.Println("Error: ",err)
	}

	c.JSON(http.StatusOK, organisms)
}

func GetOrganismById(c *gin.Context) {
	// Get id and converts it
	id := convertStringToInt(c.Param("id"))

	var err error
	var organism models.Organism

	organism, err = models.FindOrganismByID(id)

	if err != nil {
		fmt.Println("Error: ",err)
	}

	c.JSON(http.StatusOK, organism)
}

func DeleteOrganismById(c *gin.Context) {
	// Get id and converts it
	id := convertStringToInt(c.Param("id"))

	var err error
	err = models.DeleteOrganismByID(id)

	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return
	}

	message := "Organism with id " + c.Param("id") + " have been deleted."
	c.JSON(http.StatusOK, message)
}

func EditOrganismById(c *gin.Context) {
	// Get id and converts it
	id := convertStringToInt(c.Param("id"))

	var organism models.Organism
	if err := c.ShouldBindJSON(&organism); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var err error
	err = models.EditOrganismByID(&organism, id)

	if err != nil {
		c.JSON(http.StatusNotModified, "")
		return
	}

	c.JSON(http.StatusOK, organism)
}

func CreateOrganism(c *gin.Context) {
	var organism models.Organism
	if err := c.ShouldBindJSON(&organism); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err := models.CreateOrganism(&organism)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	c.JSON(http.StatusCreated, organism)
}

