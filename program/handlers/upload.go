package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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
		c.JSON(400, "file is not a video")
		return
	}

	if err := os.MkdirAll("./filesystem/video", 0770); err != nil {
		c.JSON(400, fmt.Sprintf("error: %v", err))
		return

	}

	c.SaveUploadedFile(file, "./filesystem/video/"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
