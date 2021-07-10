package routes

import (
	"choco/server/internals/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(router *gin.Engine, auth *auth.Auth) {
	router.GET("/api/auth/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok!")
	})
}
