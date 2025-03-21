package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"key": "value",
	})
}

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		log.Panic(err)
	}

	log.Println(file.Filename)
}
