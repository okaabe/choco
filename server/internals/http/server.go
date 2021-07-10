package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer(address string) error {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	return r.Run(address)
}
