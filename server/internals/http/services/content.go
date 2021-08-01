package services

import (
	"choco/server/internals/content"
	"choco/server/internals/http/inputs"
	"net/http"

	// "net/http"

	"github.com/gin-gonic/gin"
)

type ContentService struct {
	Content *content.Content
}

func (this *ContentService) CreateCommunity(c *gin.Context) {
	var (
		form inputs.CreateCommunity
		token = c.Request.Header.Get("Authorization")
	)

	if err := c.ShouldBindJSON(&form); err != nil || token == "" || form.Invalid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	community, commErr := this.Content.CreateCommunity(form.Name, form.Description, token, form.Nsfw, form.Private)

	if commErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": commErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": community,
	})
}

func (this *ContentService) CreatePost(c *gin.Context) {

}

func (this *ContentService) GetJoinedCommunities(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	members, err := this.Content.GetJoinedCommunities(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": members,
	})
}
