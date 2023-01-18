package infoshera

import (
	"github.com/gin-gonic/gin"
	"infoSfera_proxy/internal/config"
	"infoSfera_proxy/pkg/save_file"
	"net/http"
)

type reqJ struct {
	Data map[string]string `json:"data"`
}

func PostRequest(app *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestJson := reqJ{}
		if err := c.BindJSON(&requestJson); err == nil {
			if len(requestJson.Data) == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Not valid JSON"})
			} else {
				for k, v := range requestJson.Data {
					saveFileData := save_file.SaveFileData{
						IsRequest:  true,
						FileName:   k,
						StringData: v,
					}
					err = saveFileData.SaveFile(app)
					if err != nil {
						c.AbortWithError(http.StatusBadRequest, err)
					}
				}
				c.JSON(http.StatusAccepted, gin.H{"status": "success"})
			}
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	}
}
