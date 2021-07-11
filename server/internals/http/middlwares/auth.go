package middlwares

import (
	"choco/server/internals/auth"

	"github.com/gin-gonic/gin"
)

type AuthMiddlware struct {
	Auth *auth.Auth
}

func (this *AuthMiddlware) Middlware(c *gin.Context) {
	c.Next()
}