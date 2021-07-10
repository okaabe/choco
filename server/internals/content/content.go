package content

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
)

/**
- NewCommunity(name, description, token string, nsfw, private bool)
**/

type Content struct {
	Auth        *auth.Auth
	CommunityAdapter adapters.CommunityAdapter
	PostAdapter adapters.PostAdapter
}
