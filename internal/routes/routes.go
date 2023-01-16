package routes

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/handlers"
)

func Route(app *config.AppConfig) *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")

	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.Static("/images", "./static/images")
	router.Static("/files", "./static/files")

	router.GET("/", handlers.IndexHandler())
	return router
}
