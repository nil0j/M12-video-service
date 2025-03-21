package routes

import (
	// "html/template"

	"text/template"

	"github.com/gin-gonic/gin"

	"github.com/nil0j/M12-video-service/front/pkg/handlers"
)

var router = gin.Default()

func setupServer() {
	router.SetFuncMap(template.FuncMap{
		"hello": printHello,
	})
	router.LoadHTMLGlob("templates/*")
	// template.Must(template.ParseGlob())
}

func Run() {
	setupServer()

	router.GET("/", handlers.MainPage)
	router.POST("/", handlers.UploadVideo)
	router.Run("localhost:8080")
}

func printHello() string {
	return "Hell-o"
}
