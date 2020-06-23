package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/JackMaarek/spiderMail/Models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users = Models.GetAllUsers()
	c.JSON(http.StatusOK, users)
}