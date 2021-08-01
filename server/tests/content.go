package tests

import (
	"choco/server/internals/auth"
	"choco/server/internals/content"
	"choco/server/internals/models"
	"testing"
)

func testContentValidCreateCommunity(t *testing.T, content *content.Content, validToken string) *models.Community {
	community, communityErr := content.CreateCommunity("Choco tests", "A simple community to work as...", validToken, false, false)

	if communityErr != nil {
		t.Errorf("Not expected an error to create a community: %s\n", communityErr)
	}

	return community
}

func testContentInvalidCreateCommunity(t *testing.T, content *content.Content, invalidToken string) {
	_, communityErr := content.CreateCommunity("Choco tests", "A simple community to work as...", invalidToken, false, false)

	if communityErr == nil {
		t.Errorf("Not expected an error to create a community: %s\n", communityErr)
	}
}

func testContentValidJoinTheCommunity(t *testing.T, content *content.Content, token, communityId string) *models.Member {
	member, err := content.JoinTheCommunity(token, communityId)

	if err != nil {
		t.Errorf("Not expected an error to join on a community: %s", err)
	}

	return member
}

func testContentInvalidJoinTheCommunity(t *testing.T, content *content.Content, token, communityId string) {
	_, err := content.JoinTheCommunity(token, communityId)

	if err == nil {
		t.Errorf("Expected an error to try to join on a community that doesnt exists: %s", err)
	}
}

func testContentValidCreatePost(t *testing.T, content *content.Content, token, communityId, memberId, title, text string, private, nsfw bool) *models.Post {
	post, err := content.CreatePost(title, text, token, communityId, private, nsfw)

	if err != nil || post == nil {
		t.Errorf("Not expected an error to create a post: %s", err)
	}

	return post
}

func testContentInvalidCreatePost(t *testing.T, content *content.Content, token, communityId, memberId, title, text string, private, nsfw bool) {
	post, err := content.CreatePost(title, text, token, communityId, private, nsfw)

	if err == nil || post != nil {
		t.Errorf("Expected an error to try to create a post with invalid values(token, communityId, memberId...): %s", err)
	}
}

func testContentValidSearch(t *testing.T, content *content.Content, text string) ([]models.Community, []models.Post) {
	communities, posts, err := content.Search(text)

	if err != nil {
		t.Errorf("Not expected an error to search content: %s", err)
	}

	if len(communities) < 1 && len(posts) < 1 {
		t.Errorf("Expected something different as result of the search, but got an empty array of communities and posts: %s", err)
	}

	return communities, posts
}

func testContentValidGetJoinedCommunities(t *testing.T, content *content.Content, token string) []models.Member {
	members, err := content.GetJoinedCommunities(token)

	if err != nil {
		t.Errorf("Not expected an error to get the communities: %s", err)
	}

	if len(members) < 1 {
		t.Errorf("Expected more than 1 community: %s", err)
	}

	return members
}

func testContentInvalidGetJoinedCommunities(t *testing.T, content *content.Content, invalidToken string) []models.Member {
	members, err := content.GetJoinedCommunities(invalidToken)

	if err == nil {
		t.Errorf("Expected an error to get the communities: %s", err)
	}

	return members
}

func testContent(t *testing.T, auth *auth.Auth, content *content.Content) {
	_, token, registerErr := auth.Register("okaabe", "okaabe@okaabe", []byte("okaabe"))

	if registerErr != nil {
		t.Errorf("Not expected an error to register an user to test the content features")
	}

	invalidToken := "apkwpkad.kdpwoakdpawkdpakdw.dkwpoadkwd"

	var community = testContentValidCreateCommunity(t, content, token)
	testContentInvalidCreateCommunity(t, content, invalidToken)

	member := testContentValidJoinTheCommunity(t, content, token, community.ID)
	testContentInvalidJoinTheCommunity(t, content, token, "dawww.a.a.adddd.d")

	post := testContentValidCreatePost(t, content, token, community.ID, member.ID, "hello world", "hello world", false, false)
	testContentInvalidCreatePost(t, content, token, "a.a.a", member.ID, "hello world", "hello world 2", false, false)

	testContentValidSearch(t, content, post.Text)
	testContentValidGetJoinedCommunities(t, content, token)

	testContentValidGetJoinedCommunities(t, content, token)
	testContentInvalidGetJoinedCommunities(t, content, "d.d.d.d.d")
}
