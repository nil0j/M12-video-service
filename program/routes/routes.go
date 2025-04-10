package routes

import (
	"net/http"

	"github.com/nil0j/video-service/handlers"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.MainPage)
	router.POST("/", handlers.UploadVideo)
	router.StaticFS("/fs", http.Dir("./filesystem"))

	router.SetTrustedProxies(nil)
	router.Run("localhost:8080")
}
