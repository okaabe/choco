package routes

import (
	"choco/server/internals/http/middlwares"
	"choco/server/internals/http/services"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(router *gin.Engine, sessionService *services.AuthService, middlware *middlwares.AuthMiddlware) {
	router.POST("/api/auth/signin", sessionService.SignIn)
	router.POST("/api/auth/signup", sessionService.SignUp)
	router.GET("/api/auth/rewoke", middlware.Middlware, sessionService.Rewoke)
}
