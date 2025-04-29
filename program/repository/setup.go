package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/nil0j/jirafeitor/config"
)

var conn *pgx.Conn

func Setup() error {
	log.Println("Setting up database")

	var err error
	conn, err = pgx.Connect(context.Background(), config.Data.Db_url)
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())
	return nil
}
