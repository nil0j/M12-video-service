package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var Data = appData{}

type appData struct {
	Filesystem string
	Db_url     string
	Secret     string
}

func Setup() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("Env file not found, skipping...")
	}

	if os.Getenv("DB_URL") == "" {
		return errors.New("Environment variable DB_URL not set")
	}

	if os.Getenv("SECRET") == "" {
		return errors.New("Environment variable SECRET not set")
	}

	Data = appData{
		Filesystem: "./filesystem",
		Db_url:     os.Getenv("DB_URL"),
	}

	return nil
}
