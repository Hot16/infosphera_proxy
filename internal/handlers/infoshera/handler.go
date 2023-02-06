package infoshera

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/internal/models"
	"infoSfera_proxy/pkg/send_request"
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

						cred := send_request.NewCred("infoshera")
						cred.Id = k
						cred.PostFields = []byte(v)

						config.App.SendRequest <- cred
					}(k, v)
				}
				c.JSON(http.StatusAccepted, gin.H{"status": "success"})
			}
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
