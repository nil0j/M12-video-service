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
	baseGroup.GET("/video", handlers.VideoPage)
	baseGroup.GET("/video/:id", handlers.GetVideo)

	baseGroup.POST("/", handlers.Login)

	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
