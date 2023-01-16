package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Start",
		})
	}
}
