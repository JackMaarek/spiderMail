package services

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/JackMaarek/spiderMail/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateTokenExpireDate() (time.Duration, error) {
	var err error
	var key = config.GetDotEnvVariable("ACCESS_EXPIRES")
	var durationInDays int64
	var expireDate time.Duration
	durationInDays, err = strconv.ParseInt(key, 10, 32)
	if err != nil {
		return 1, err
	}
	expireDate = time.Hour * 24 * time.Duration(durationInDays)
	return expireDate, nil
}

func CreateToken(userId uint64, userAdmin bool) (string, error) {
	var err error
	//Creating Access Token
	var expireDate time.Duration
	expireDate, err = CreateTokenExpireDate()
	atClaims := jwt.MapClaims{}
	atClaims["revoked"] = false
	atClaims["admin"] = userAdmin
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(expireDate).UTC()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.GetDotEnvVariable("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func VerifyToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.GetDotEnvVariable("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(c *gin.Context) error {
	token, err := VerifyToken(c)
	if err != nil {
		return err
	}
	fmt.Println(token.Claims.(jwt.Claims))
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
