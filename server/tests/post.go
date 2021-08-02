package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/models"
	"testing"
)

func testPostAdapterCreate(t *testing.T, adapter adapters.PostAdapter, object *models.Post) *models.Post {
	err := adapter.Add(object)

	if err != nil {
		t.Errorf("Not expected an error to create a post in the database: %s", err)
	}

	return object
}

func testPostAdapterGet(t *testing.T, adapter adapters.PostAdapter, id string) *models.Post {
	post, err := adapter.Get(id)

	if err != nil || post == nil {
		t.Errorf("Not expected an error to get a post that should exists in the database: %s", err)
	}

	return post
}

func testPostAdapterInvalidGet(t *testing.T, adapter adapters.PostAdapter, id string) {
	post, err := adapter.Get(id)

	if err == nil || post != nil {
		t.Errorf("Expected an error to get a post that shouldn't exists in the database: %s", err)
	}
}

func testPostAdapterCommunity(t *testing.T, adapter adapters.PostAdapter, community_id string) []models.Post {
	posts, err := adapter.Community(community_id)

	if err != nil {
		t.Errorf("Not expected an error to get the posts of a community: %s", err)
	}

	return posts
}

func testPostAdapter(t *testing.T, adapter adapters.PostAdapter) {
	community, commErr := models.NewCommunity("choco", "choco cool", "kdpoakwdop", false)

	if commErr != nil {
		t.Errorf("Not expected an error to create a community model to test the post adapter: %s", commErr)
	}

	post, postErr := models.NewPost("A cool title about chocolate", "A big text that describes how chocolate is amazing", "kdpoakwdop", community.ID, false)

	if postErr != nil {
		t.Errorf("Not expected an error to create a post model to test the post adapter: %s", postErr)
	}

	testPostAdapterCreate(t, adapter, post)
	testPostAdapterGet(t, adapter, post.ID)
	testPostAdapterInvalidGet(t, adapter, "dkapapapapapapawekekekekek")
	testPostAdapterCommunity(t, adapter, community.ID)
}
