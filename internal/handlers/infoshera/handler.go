package infoshera

import (
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/pkg/save_file"
	"infoSfera_proxy/pkg/send_request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reqJ struct {
	Data map[string]string `json:"data"`
}

func PostRequest(app *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestJson := reqJ{}
		if err := c.BindJSON(&requestJson); err == nil {
			if len(requestJson.Data) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid JSON"})
			} else {
				for k, v := range requestJson.Data {
					saveFileData := save_file.SaveFileData{
						IsRequest:  true,
						FileName:   k,
						StringData: v,
					}

					go saveFileData.SaveFile(app)

					credentials := send_request.Credentials{
						BaseUrl:   app.Env.GetString("external.weatherapi-weather.baseUrl"),
						Method:    app.Env.GetString("external.weatherapi-weather.method"),
						Headers:   make(map[string]string),
						GetParams: make(map[string]string),
					}
					for k, v := range app.Env.GetStringMapString("external.weatherapi-weather.headers") {
						credentials.Headers[k] = v
					}
					for k, v := range app.Env.GetStringMapString("external.weatherapi-weather.query_params") {
						credentials.GetParams[k] = v
					}
					credentials.GetParams["q"] = "Podgorica"
					go credentials.SendRequest()
				}
				c.JSON(http.StatusAccepted, gin.H{"status": "success"})
			}
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
