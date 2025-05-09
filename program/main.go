package main

import (
	"log"
	"os"

	"github.com/nil0j/jirafeitor/config"
	"github.com/nil0j/jirafeitor/repository"
	"github.com/nil0j/jirafeitor/routes"
)

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
