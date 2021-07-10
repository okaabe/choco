package routes

import (
	"choco/server/internals/auth"
	"choco/server/internals/content"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, auth *auth.Auth, content *content.Content) error {
	registerAuthRoutes(router, auth)
	registerContentRoutes(router, content)

	return nil
}
