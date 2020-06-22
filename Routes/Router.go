package Routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"spiderMail/Controllers"
)

var (
	router = gin.Default()
)

func SetupRouter(){
	router.POST("/login", Controllers.Login)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	log.Fatal(router.Run(":8080"))
}
