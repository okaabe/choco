package content

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
	"choco/server/internals/models"
	"errors"
)

type Content struct {
	Auth             *auth.Auth
	MemberAdapter    adapters.MemberAdapter
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
		return nil, errors.New("Couldn't create the community")
	}

	member, memberErr := models.NewMember(user.ID, community.ID)

	if memberErr != nil {
		return nil, errors.New("Couldn't create the community")
	}

	communityAdapterErr := this.CommunityAdapter.Add(community)

	if communityAdapterErr != nil {
		return nil, errors.New("Couldn't save the community")
	}

	memberAdapterErr := this.MemberAdapter.Add(member)

	if memberAdapterErr != nil {
		return nil, errors.New("COuldn't create the community")
	}

	return community, nil
}
