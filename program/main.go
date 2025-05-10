package main

import (
	"log"
	"os"

	"github.com/nil0j/jirafeitor/config"
	_ "github.com/nil0j/jirafeitor/docs" // Add this line
	"github.com/nil0j/jirafeitor/repository"
	"github.com/nil0j/jirafeitor/routes"
)

// @title Jirafeitor API
// @version	1.0
// @description A video service made using gin-gonic

// @securityDefinitions.apiKey JWT
// @in header
// @name token

// @host localhost:8080
// @BasePath /api
func main() {
	if err := config.Setup(); err != nil {
		log.Panic(err)
	}

	setupFilesystem()

	err := repository.Setup()
	if err != nil {
		log.Panic(err)
	}

	routes.Run()
}

func setupFilesystem() {
	err := os.MkdirAll("./filesystem", 0770)
	if err != nil {
		log.Panic(err)
	}
}
