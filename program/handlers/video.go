package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nil0j/jirafeitor/repository"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "video", gin.H{})
}

func GetVideo(c *gin.Context) {
	// c.File("./ignore/Billie Jean (low quality).mp4")
	videos, err := repository.GetAllVideos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, videos)
}
