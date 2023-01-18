package routes

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/handlers"
	"infoSfera_proxy/internal/handlers/infoshera"
)

func Route(app *config.AppConfig) *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")

	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.Static("/images", "./static/images")
	router.Static("/files", "./static/files")

	router.GET("/", handlers.IndexHandler())

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/req_infoshera", infoshera.PostRequest(app))
	}
	return router
}
