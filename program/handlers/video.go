package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VideoPage(c *gin.Context) {
	c.HTML(http.StatusOK, "video", gin.H{})
}

func GetVideo(c *gin.Context) {
	c.File("./ignore/Billie Jean (low quality).mp4")
}
