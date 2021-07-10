package http

import (
	"choco/server/internals/auth"
	"choco/server/internals/content"
	"choco/server/internals/http/routes"

	"github.com/gin-gonic/gin"
)

func RunServer(address string, auth *auth.Auth, content *content.Content) error {
	r := gin.Default()

	routes.Routes(r, auth, content)

	return r.Run(address)
}
