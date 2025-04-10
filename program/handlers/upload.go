package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		log.Println("file is not a video")
		c.JSON(400, "file is not a video")
		return
	}

	if err := os.Mkdir("./filesystem/video", 0770); err != nil {
		c.JSON(400, "file is not a video")
		return

	}

	c.SaveUploadedFile(file, "./filesystem/video/"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
