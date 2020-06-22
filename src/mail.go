package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func main() {
	router.POST("/login", Login)
	log.Fatal(router.Run(":8080"))
}

func Login(c *gin.Context) {
	println("jooj")
}