package midlewares

import (
	"fmt"
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckAuthorization(c *gin.Context) {
	var err error
	var token *models.Token
	token, err = models.FindTokenByToken(services.ExtractToken(c))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	var user *models.User
	user, err = models.FindUserByID(token.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	fmt.Println(&user)
	err = services.TokenValid(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.Next()
}
