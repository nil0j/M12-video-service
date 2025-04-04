package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadVideo(c *gin.Context) {
	file, err := c.FormFile("video")
	if err != nil {
		log.Println(err)
		c.JSON(500, "error")
		return
	}

	c.SaveUploadedFile(file, "./filesystem/"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
