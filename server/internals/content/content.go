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

func (this *Content) JoinTheCommunity(token string, communityId string) (*models.Member, error) {
	user, rewokeErr := this.Auth.Rewoke(token)

	if rewokeErr != nil {
		return nil, rewokeErr
	}

	comm, commErr := this.CommunityAdapter.Get(communityId)

	if commErr != nil {
		return nil, errors.New("COuldn't find a community with this id")
	}

	member, memberErr := models.NewMember(user.ID, comm.ID)

	if memberErr != nil {
		return nil, errors.New("Couldn't join on the community")
	}

	adapterErr := this.MemberAdapter.Add(member)

	if adapterErr != nil {
		return nil, errors.New("Couldn't join on the community")
	}

	return member, nil
}

func (this *Content) CreatePost(title, text, token, communityId string, private, nsfw bool) (*models.Post, error) {
	user, rewokeErr := this.Auth.Rewoke(token)

	if rewokeErr != nil {
		return nil, rewokeErr
	}

	community, communityErr := this.CommunityAdapter.Get(communityId)

	if communityErr != nil {
		return nil, errors.New("Couldn't find the community with this id")
	}

	member, memberErr := this.MemberAdapter.MemberInTheCommunity(community.ID, user.ID)

	if memberErr != nil {
		return nil, errors.New("You don't have permission to create a post on this guild, since that you are not a member of it")
	}

	post, postErr := models.NewPost(title, text, member.ID, community.ID, private, nsfw)

	if postErr != nil {
		return nil, errors.New("Couldn't create the post")
	}

	adapterErr := this.PostAdapter.Add(post)

	if adapterErr != nil {
		return nil, adapterErr
	}

	return post, nil
}

func (this *Content) Search(text string) ([]models.Community, []models.Post, error) {
	communities, communitiesErr := this.CommunityAdapter.Search(text)

	if communitiesErr != nil {
		return nil, nil, errors.New("Couldn't find anything")
	}

	posts, postsErr := this.PostAdapter.Search(text)

	if postsErr != nil {
		return nil, nil, errors.New("Couldn't find anything")
	}

	return communities, posts, nil
}

func (this *Content) GetCommunities(token string) ([]models.Community, error) {
	return nil, nil
}
