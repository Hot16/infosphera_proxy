package infoshera

import (
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/models"
	"infoSfera_proxy/pkg/send_request"
	"net/http"

	"github.com/gin-gonic/gin"
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

						credentials := send_request.Credentials{
							BaseUrl:   config.App.Env.GetString("external.weatherapi-weather.baseUrl"),
							Method:    config.App.Env.GetString("external.weatherapi-weather.method"),
							Headers:   make(map[string]string),
							GetParams: make(map[string]string),
						}
						for k, v := range config.App.Env.GetStringMapString("external.weatherapi-weather.headers") {
							credentials.Headers[k] = v
						}
						for k, v := range config.App.Env.GetStringMapString("external.weatherapi-weather.query_params") {
							credentials.GetParams[k] = v
						}
						credentials.GetParams["q"] = "Podgorica"
						go credentials.SendRequest()
					}(k, v)
				}
				c.JSON(http.StatusAccepted, gin.H{"status": "success"})
			}
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
