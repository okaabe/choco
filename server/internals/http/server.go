package http

import (
	"choco/server/internals/content"
	"choco/server/internals/http/routes"
	"choco/server/internals/session"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func RunServer(address string, auth *session.SessionUseCase, content *content.Content) error {
	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	routes.Routes(router, auth, content)

	return router.Run(address)
}
