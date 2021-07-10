package routes

import (
	"choco/server/internals/content"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerContentRoutes(router *gin.Engine, content *content.Content) {
	router.GET("/api/content/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Ok!")
	})
}
