package services

import (
	"choco/server/internals/http/inputs"
	"choco/server/internals/usecase/session"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthService struct {
	Session *session.SessionUseCase
}

func (this *AuthService) SignUp(c *gin.Context) {
	var signup inputs.SignUp

	if err := c.ShouldBindJSON(&signup); err != nil || signup.Email == "" || signup.Password == "" || signup.Username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	user, token, err := this.Session.Register(signup.Username, signup.Email, []byte(signup.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"jwt":        token,
		},
	})
}

func (this *AuthService) SignIn(c *gin.Context) {
	var signin inputs.SignIn

	if err := c.ShouldBindJSON(&signin); err != nil || signin.Email == "" || signin.Password == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	user, token, err := this.Session.Authenticate(signin.Email, []byte(signin.Password))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
			"jwt":        token,
		},
	})
	return
}

func (this *AuthService) Rewoke(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	user, err := this.Session.Rewoke(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}
