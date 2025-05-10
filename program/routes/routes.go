package routes

import (
	"github.com/nil0j/jirafeitor/handlers"
	"github.com/nil0j/jirafeitor/middlewares"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	baseGroup := router.Group("api")

	baseGroup.POST("login", handlers.Login)
	baseGroup.POST("register", middlewares.Ultrapass(), handlers.Register)

	{
		uploadGroup := router.Group("upload", middlewares.JWT())
		uploadGroup.POST("", handlers.UploadVideo)
	}

	baseGroup.GET("videos", handlers.GetRecentVideos)
	{
		videoGroup := baseGroup.Group("video")
		videoGroup.GET("source/:id", handlers.GetVideo)
		videoGroup.GET("info/:id", handlers.GetVideoInfo)
		videoGroup.GET("thumb/:id", handlers.GetVideoThumbnail)
	}

	router.SetTrustedProxies(nil)
	gin.SetMode(gin.ReleaseMode)
	router.Run("localhost:8080")
}
