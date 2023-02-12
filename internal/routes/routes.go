package routes

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/handlers"
)

func Route() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("./templates/*")

	router.Static("/css", "./static/css")
	router.Static("/js", "./static/js")
	router.Static("/images", "./static/images")
	router.Static("/files", "./static/files")

	router.GET("/", handlers.IndexHandler())

	apiGroup := router.Group("/api")
	{
		apiGroup.GET("/req_infosphera", handlers.GetApiReq())
		apiGroup.POST("/req_infosphera", handlers.PostRequest())
	}
	return router
}
