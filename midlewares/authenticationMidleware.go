package midlewares

import (
	"errors"
	"fmt"
	"github.com/JackMaarek/spiderMail/services"
	"github.com/gin-gonic/gin"
)

func CheckAuthorization(c *gin.Context) error {
	var err error
	err = services.TokenValid(c)
	fmt.Println(err)
	if err != nil {
		return errors.New("Not Authorized")

	}
	return nil
}