package main

import (
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/routes"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"log"
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
	routes.SetupRouter(router)

	log.Fatal(router.Run(":8080"))
}
