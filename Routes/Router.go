package Routes

import (
	"github.com/JackMaarek/spiderMail/Controllers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	router = gin.Default()
)

func SetupRouter(){
	router.POST("/login", Controllers.Login)
	router.GET("/users", Controllers.GetUsers)
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	log.Fatal(router.Run(":8080"))
}
