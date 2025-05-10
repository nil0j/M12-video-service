package handlers

import (
	"encoding/json"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/nil0j/jirafeitor/models/postgres"
	"github.com/nil0j/jirafeitor/models/responses"
	"github.com/nil0j/jirafeitor/repository"
)

func Register(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	user := postgres.PostgresUserPost{}
	if err := json.Unmarshal(jsonData, &user); err != nil {
		responses.HandleError(c, err)
		return
	}

	token, err := repository.CreateUser(user)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleJWTLogin(c, user.Username, token)
}

func Login(c *gin.Context) {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	user := postgres.PostgresUserPost{}
	if err := json.Unmarshal(jsonData, &user); err != nil {
		responses.HandleError(c, err)
		return
	}

	token, err := repository.LogIn(user)
	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.HandleJWTLogin(c, user.Username, token)
}
