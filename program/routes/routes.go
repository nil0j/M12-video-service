package routes

import (
	"github.com/nil0j/jirafeitor/handlers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*/*")

	router.GET("/", handlers.HomePage)
	router.GET("/upload", handlers.UploadPage)
	router.POST("/upload", handlers.UploadVideo)
	router.GET("/video", handlers.VideoPage)
	router.GET("/video/:id", handlers.GetVideo)

	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
