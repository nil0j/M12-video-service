package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Data = appData{}

type appData struct {
	Filesystem string
	Db_url     string
}

func Setup() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env file not found, skipping...")
	}

	Data = appData{
		Filesystem: "./filesystem",
		Db_url:     os.Getenv("DB_URL"),
	}
}
