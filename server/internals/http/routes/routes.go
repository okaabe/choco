package routes

import (
	"choco/server/internals/content"
	"choco/server/internals/http/middlwares"
	"choco/server/internals/http/services"
	"choco/server/internals/session"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, sessionUseCase *session.SessionUseCase, content *content.Content) error {
	authMiddlware := &middlwares.AuthMiddlware{
		Session: sessionUseCase,
	}

	registerAuthRoutes(router, &services.AuthService{
		Session: sessionUseCase,
	}, authMiddlware)

	registerContentRoutes(router, &services.ContentService{
		Content: content,
	}, authMiddlware)

	return nil
}
