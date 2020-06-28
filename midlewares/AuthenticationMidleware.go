package midlewares

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuthorization(c *gin.Context) {
	var err error
	err = services.TokenValid(c)
	fmt.Println("error: ", err)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.Next()
}
