package routes

import (
	"choco/server/internals/auth"
	"choco/server/internals/content"
	"choco/server/internals/http/middlwares"
	"choco/server/internals/http/services"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, auth *auth.Auth, content *content.Content) error {
	authMiddlware := &middlwares.AuthMiddlware{
		Auth: auth,
	}

	registerAuthRoutes(router, &services.AuthService{
		Auth: auth,
	}, authMiddlware)

	return nil
}
