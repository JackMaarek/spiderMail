package routes

import (
	"github.com/JackMaarek/spiderMail/controllers"
	"github.com/JackMaarek/spiderMail/midlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/registration", controllers.Registration)

	authorized := router.Group("/")
	authorized.Use(midlewares.CheckAuthorization)
	{
		authorized.PUT("/users/:id", controllers.UpdateUser)
		authorized.DELETE("/users/:id", controllers.DeleteUser)
		authorized.GET("/users/refresh_token", controllers.RefreshToken)
		authorized.GET("/organisms/:id/users", controllers.GetUsersByOrganism)
		authorized.GET("/organisms", controllers.GetOrganisms)
		authorized.GET("/organisms/:id", controllers.GetOrganismById)
		authorized.POST("/organisms", controllers.CreateOrganism)
		authorized.DELETE("/organisms/:id", controllers.DeleteOrganismById)
		authorized.PUT("/organisms/:id", controllers.EditOrganismById)

		authorized.GET("/campaigns", controllers.GetCampaigns)
		authorized.GET("/campaigns/:id", controllers.GetCampaignById)
		authorized.PUT("/campaigns/:id", controllers.EditCampaignById)
		authorized.DELETE("/campaigns/:id", controllers.DeleteCampaignById)
		authorized.POST("/campaigns", controllers.CreateCampaign)
		authorized.GET("/organisms/:id/campaigns", controllers.GetCampaignsByOrganismId)

		authorized.GET("/groups", controllers.GetRecipientList)
		authorized.GET("/groups/:id", controllers.GetRecipientsListById)
		authorized.POST("/groups", controllers.CreateRecipientsList)
		authorized.PUT("/groups/:id", controllers.EditRecipientsListById)
		authorized.DELETE("/groups/:id", controllers.DeleteRecipientsListById)

	}
}
