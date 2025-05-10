package routes

import (
	"github.com/nil0j/jirafeitor/handlers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*/*")
	baseGroup := router.Group("api")

	baseGroup.GET("/", handlers.HomePage)
	baseGroup.GET("/upload", handlers.UploadPage)
	baseGroup.POST("/upload", handlers.UploadVideo)

	videoGroup := baseGroup.Group("video")
	baseGroup.GET("video", handlers.VideoPage)

	videoGroup.GET("source/:id", handlers.GetVideo)
	videoGroup.GET("info/:id", handlers.GetVideoInfo)
	videoGroup.GET("thumb/:id", handlers.GetVideoThumbnail)

	baseGroup.POST("/", handlers.Login)

	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
