package handlers

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/handlers/infoshera"
	"infoSfera_proxy/internal/models"
	"net/http"
)

type reqJ struct {
	Data map[string]string `json:"data"`
}

func PostRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestJson := reqJ{}
		if err := c.BindJSON(&requestJson); err == nil {
			if len(requestJson.Data) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON"})
			} else {
				for k, v := range requestJson.Data {
					go func(k string, v string) {
						saveFileData := models.SaveFileData{
							IsRequest:  true,
							FileName:   k,
							StringData: v,
						}
						config.App.SaveFileChan <- saveFileData

						go infoshera.InfoSpheraRequest(k, v)
					}(k, v)
				}
				c.JSON(http.StatusAccepted, gin.H{"status": "success"})
			}
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Start",
		})
	}
}
