package main

import (
	"log"
	"os"

	"github.com/nil0j/video-service/routes"
)

func setup() {
	err := os.MkdirAll("./filesystem", 0770)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	setup()
	routes.Run()
}
