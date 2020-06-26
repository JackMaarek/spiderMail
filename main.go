package main

import (
	"github.com/JackMaarek/spiderMail/models"
	"github.com/JackMaarek/spiderMail/routes"
	//"github.com/JackMaarek/spiderMail/services/rabbitmq/producer"
	"github.com/JackMaarek/spiderMail/rabbitmq/consummer"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"log"
	"time"
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
	//producer.SendToRabbit()
	consummer.ReceiveToRabbit()

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

	log.Fatal(router.Run(":8080"))
}
