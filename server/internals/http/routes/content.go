package routes

import (
	"choco/server/internals/http/middlwares"
	"choco/server/internals/http/services"

	"github.com/gin-gonic/gin"
)

func registerContentRoutes(router *gin.Engine, content *services.ContentHttpService, middlware *middlwares.AuthMiddlware) {
	router.POST("/api/content/community", middlware.Middlware, content.CreateCommunity)

	router.GET("/api/content/user/communities", middlware.Middlware, content.GetJoinedCommunities)
	router.POST("/api/content/user/community/", middlware.Middlware, content.JoinCommunity)

	router.GET("/api/content/", content.Search)
	router.GET("/api/content/community/:name", content.GetCommunity)

	router.POST("/api/content/community/:name", middlware.Middlware, content.CreatePost)
	router.GET("/api/content/community/:name/posts", content.GetPosts)
}
