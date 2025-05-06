package repository

import (
	"github.com/nil0j/jirafeitor/models/postgres"

	"context"
)

func CreateUser(username string, password string) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO jirafeitor.users (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(username string) (*postgres.PostgresUser, error) {
	var user postgres.PostgresUser
	err := conn.QueryRow(context.Background(), "SELECT id, username, password FROM jirafeitor.users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
