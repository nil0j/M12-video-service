package repository

import (
	"context"

	"github.com/nil0j/jirafeitor/models/errorresponses"
	"github.com/nil0j/jirafeitor/models/postgres"
	"github.com/nil0j/jirafeitor/utils/auth"
)

func CreateUser(user postgres.PostgresUserPost) (string, error) {
	var tmpUsername string
	err := conn.QueryRow(context.Background(), "SELECT username FROM users WHERE username = $1", user.Username).Scan(&tmpUsername)
	if err == nil {
		return "", errorresponses.UserNotFound
	}

	passwordHash, err := auth.Hash(user.Password)
	if err != nil {
		return "", err
	}

	if _, err := conn.Exec(context.Background(), "INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, passwordHash); err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}

func LogIn(user postgres.PostgresUserPost) (string, error) {
	var response postgres.PostgresUser
	err := conn.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", user.Username).Scan(&response.ID, &response.Username, &response.Password)
	if err != nil {
		return "", err
	}

	if auth.CheckPasswordHash(user.Password, response.Password) {
		token, err := auth.GenerateToken(user.Username)
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", errorresponses.PasswordDontMatch
}
