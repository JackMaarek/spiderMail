package midlewares

import (
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
	var _ *models.User
	_, err = models.FindUserByID(token.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}
	err = services.TokenValid(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.Next()
}
