package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "video", gin.H{})
}

func GetVideo(c *gin.Context) {
	folderPath := "./filesystem/" + c.Param("id")
	files, err := os.ReadDir(folderPath)
	if err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}

	if len(files) == 0 {
		c.JSON(400, gin.H{"message": "file not found"})
		return
	}

	filePath := filepath.Join(folderPath, files[0].Name())

	c.File(filePath)
}
