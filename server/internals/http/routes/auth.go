package routes

import (
	"choco/server/internals/http/middlwares"
	"choco/server/internals/http/services"

	"github.com/gin-gonic/gin"
)

func registerAuthRoutes(router *gin.Engine, auth *services.AuthService, middlware *middlwares.AuthMiddlware) {
	router.POST("/api/auth/signin", auth.SignIn)
	router.POST("/api/auth/signup", auth.SignUp)
	router.GET("/api/auth/rewoke", middlware.Middlware, auth.Rewoke)
}
