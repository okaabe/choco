package content

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
	"choco/server/internals/models"
)

type Content struct {
	Auth             *auth.Auth
	MemberAdapter    adapters.MemberAdapter
	CommunityAdapter adapters.CommunityAdapter
	PostAdapter      adapters.PostAdapter
}

func (this *Content) CreateCommunity(name, description, token string, nsfw, private bool) (*models.Community, error) {

	return nil, nil
}
