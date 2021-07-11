package services

import (
	"choco/server/internals/auth"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	Auth *auth.Auth
}

func (this *AuthService) SignUp(c *gin.Context) {
	c.JSON(200, "message")
}

func (this *AuthService) SignIn(c *gin.Context) {
	c.JSON(200, "message")
}

func (this *AuthService) Rewoke(c *gin.Context) {
	c.JSON(200, "message")
}
