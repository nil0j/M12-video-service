package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nil0j/jirafeitor/repository"
)

func UploadPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload", gin.H{})
}

func UploadVideo(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	_, _ = title, description

	file, err := c.FormFile("video")
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("%v", err),
		})
		return
	}

	if err := os.MkdirAll("./filesystem/video", 0770); err != nil {
		c.JSON(400, fmt.Sprintf("error: %v", err))
		return

	}

	id, err := repository.CreateVideo(title, description)
	if err != nil {
		c.JSON(400, fmt.Sprintf("error: %v", err))
		return
	}

	c.SaveUploadedFile(file, "./filesystem/"+fmt.Sprintf("%d", id)+"/"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
