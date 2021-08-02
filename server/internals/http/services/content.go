package services

import (
	"choco/server/internals/http/binds"
	"choco/server/internals/usecase/content"
	"net/http"

	// "net/http"

	"github.com/gin-gonic/gin"
)

type ContentHttpService struct {
	Content *content.Content
}

func (this *ContentHttpService) CreateCommunity(c *gin.Context) {
	var (
		form  binds.CreateCommunity
		token = c.Request.Header.Get("Authorization")
	)

	if err := c.ShouldBindJSON(&form); err != nil || token == "" || form.Invalid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	community, commErr := this.Content.CreateCommunity(form.Name, form.Description, token, form.Nsfw)

	if commErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": commErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": community,
	})
}

func (this *ContentHttpService) CreatePost(c *gin.Context) {
	var (
		form          binds.CreatePost
		token         = c.Request.Header.Get("Authorization")
		communityName = c.Param("name")
	)

	if err := c.ShouldBindJSON(&form); err != nil || form.Invalid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	post, postErr := this.Content.CreatePost(form.Title, form.Text, token, communityName, form.Nsfw)

	if postErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": postErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": post,
	})
}

func (this *ContentHttpService) JoinCommunity(c *gin.Context) {
	var (
		form  binds.JoinCommunity
		token = c.Request.Header.Get("Authorization")
	)

	if err := c.ShouldBindJSON(&form); err != nil || form.Invalid() || token == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",
		})
		return
	}

	member, joinCmmErr := this.Content.JoinTheCommunity(token, form.CommunityName)

	if joinCmmErr != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": joinCmmErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": member,
	})
}

func (this *ContentHttpService) GetJoinedCommunities(c *gin.Context) {
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

func (this *ContentHttpService) Search(c *gin.Context) {
	search_query := c.Request.URL.Query().Get("search_query")

	communities, posts, err := this.Content.Search(search_query)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": gin.H{
			"communities": communities,
			"posts":       posts,
		},
	})
}

func (this *ContentHttpService) GetCommunity(c *gin.Context) {
	name := c.Param("name")

	community, communityErr := this.Content.GetCommunity(name)

	if communityErr != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"err": "Not found",
		})
		return
	}

	c.JSON(http.StatusFound, gin.H{
		"data": community,
	})
}

func (this *ContentHttpService) GetPosts(c *gin.Context) {
	
}