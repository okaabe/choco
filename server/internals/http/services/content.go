package services

import (
	"choco/server/internals/content"
	"net/http"
	// "net/http"

	"github.com/gin-gonic/gin"
)

type ContentService struct {
	Content *content.Content
}

func (this *ContentService) CreateCommunity(c *gin.Context) {

}

func (this *ContentService) CreatePost(c *gin.Context) {

}

func (this *ContentService) GetJoinedCommunities(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	_, err := this.Content.GetJoinedCommunities(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
	}

	c.JSON(201, token);
}
