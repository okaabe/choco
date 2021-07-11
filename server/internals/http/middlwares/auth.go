package middlwares

import (
	"choco/server/internals/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddlware struct {
	Auth *auth.Auth
}

func (this *AuthMiddlware) Middlware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	_, err := this.Auth.Rewoke(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid token",
		})
		return
	}

	c.Next()
}
