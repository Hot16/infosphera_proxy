package infoshera

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/pkg/save_file"
	"log"
	"net/http"
)

type reqJ struct {
	Data map[string]string `json:"data"`
}

func PostRequest(app config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestJson := reqJ{}
		if err := c.BindJSON(&requestJson); err == nil {
			for k, v := range requestJson.Data {
				saveFileData := save_file.FileData{
					IsRequest:  true,
					FileName:   k,
					StringData: v,
				}
				err = save_file.SaveFile(app, saveFileData)
				if err != nil {
					log.Println(err)
				}
			}
			c.JSON(http.StatusAccepted, gin.H{"status": "success"})
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
