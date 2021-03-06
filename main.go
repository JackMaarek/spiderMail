package main

import (
	"log"
	"time"

	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/routes"
	"github.com/JackMaarek/spiderMail/routines"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type config struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	// Database initialization
	models.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
	models.MakeMigrations()
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "Authorization",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.SetupRouter(router)

	// Check every 2 minutes if there are campaigns to send
	go routines.CheckForCampaignsToSend(2)

	log.Fatal(router.Run(":8081"))
}
