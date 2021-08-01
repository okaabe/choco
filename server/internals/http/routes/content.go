package routes

import (
	"choco/server/internals/http/middlwares"
	"choco/server/internals/http/services"

	"github.com/gin-gonic/gin"
)

func registerContentRoutes(router *gin.Engine, content *services.ContentService, middlware *middlwares.AuthMiddlware) {
	router.POST("/api/content/community", middlware.Middlware, content.CreateCommunity)
	router.GET("/api/content/user/communities", middlware.Middlware, content.GetJoinedCommunities)
}