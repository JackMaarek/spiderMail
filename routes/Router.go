package routes

import (
	"github.com/JackMaarek/spiderMail/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(router *gin.Engine){
	router.POST("/login", controllers.Login)
	router.POST("/registration", controllers.Registration)
	router.PUT("/user/update/:id", controllers.UpdateUser)
	router.DELETE("/user/delete/:id", controllers.DeleteUser)

	router.GET("/organisms", controllers.GetOrganisms)
	router.GET("/organisms/:id", controllers.GetOrganismById)
	router.POST("/organisms", controllers.CreateOrganism)
	router.DELETE("/organisms/:id", controllers.DeleteOrganismById)
	router.PUT("/organisms/:id", controllers.EditOrganismById)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
}