package content

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
	"choco/server/internals/models"
)

type Content struct {
	Auth *auth.Auth
	UserAdapter adapters.UserAdapter
	PostAdapter adapters.PostAdapter
}

func (this *Content) NewCommunity(title, description, creator_id string, private, nsfw bool) (*models.Community, error) {
	return nil, nil
}

func (this *Content) NewPost(title, text, community_id, creator_id string, private, nsfw bool) (*models.Post, error) {
	return nil, nil
}
