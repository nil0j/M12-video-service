package routes

import (
	"github.com/nil0j/video-service/handlers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*/*")

	router.GET("/", handlers.HomePage)
	router.GET("/video", handlers.VideoPage)
	router.GET("/video/:id", handlers.GetVideo)
	router.GET("/upload", handlers.UploadPage)
	router.POST("/upload", handlers.UploadVideo)

	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
