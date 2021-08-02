package middlwares

import (
	"choco/server/internals/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthMiddlware struct {
	Session *session.SessionUseCase
}

func (this *AuthMiddlware) Middlware(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	_, err := this.Session.Rewoke(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": "Invalid token",
		})
		return
	}

	c.Next()
}
