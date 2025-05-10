package handlers

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nil0j/jirafeitor/config"
	"github.com/nil0j/jirafeitor/models/responses"
	"github.com/nil0j/jirafeitor/repository"
	"github.com/nil0j/jirafeitor/utils/thumbnail"
)

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

	id, err := repository.CreateVideo(title, description)
	if err != nil {
		c.JSON(400, fmt.Sprintf("error: %v", err))
		return
	}

	videoOutFolderPath := config.Data.Filesystem + fmt.Sprintf("%d", id) + "/"
	if err := os.MkdirAll(videoOutFolderPath, 0755); err != nil {
		c.JSON(400, fmt.Sprintf("error: %v", err))
		return
	}

	videoOutPath := videoOutFolderPath + file.Filename
	c.SaveUploadedFile(file, videoOutPath)

	thumbnail.Gen(videoOutPath, videoOutFolderPath)
	responses.HandleSuccess(c, "File uploaded correctly!")
}
