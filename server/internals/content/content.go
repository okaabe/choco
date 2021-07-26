package content

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
	"choco/server/internals/models"
	"errors"
)

type Content struct {
	Auth             *auth.Auth
	UserAdapter      adapters.UserAdapter
	CommunityAdapter adapters.CommunityAdapter
	PostAdapter      adapters.PostAdapter
}

func (this *Content) CreateCommunity(name, description, token string, nsfw, private bool) (*models.Community, error) {
	user, rewokeErr := this.Auth.Rewoke(token)

	if rewokeErr != nil {
		return nil, rewokeErr
	}

	community, communityErr := models.NewCommunity(name, description, user.ID, private, nsfw)

	if communityErr != nil {
		return nil, communityErr
	}

	communityAdapterErr := this.CommunityAdapter.Add(community)

	if communityAdapterErr != nil {
		return nil, errors.New("Couldn't save the community")
	}

	user.JoinedCommunities = append(user.JoinedCommunities, community.ID)

	userAdapterErr := this.UserAdapter.UpdateCommunities(user.ID, user.JoinedCommunities)

	if userAdapterErr != nil {
		return nil, errors.New("Couldn't join on the community")
	}

	return community, nil
}
