package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"key": "value",
	})
}
