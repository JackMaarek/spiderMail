package main

import (
	"github.com/JackMaarek/spiderMail/Models"
	"github.com/JackMaarek/spiderMail/Routes"
)

func main() {
	Routes.SetupRouter()
	Models.MakeMigrations()
}