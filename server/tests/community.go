package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/models"
	"testing"
)

func testCommunityAdapterCreate(t *testing.T, adapter adapters.CommunityAdapter, object *models.Community) *models.Community {
	err := adapter.Add(object)

	if err != nil {
		t.Errorf("Not expected an error to create a community in the database: %s", err)
	}

	return object
}

func testCommunityAdapterGet(t *testing.T, adapter adapters.CommunityAdapter, id string) *models.Community {
	community, err := adapter.Get(id)

	if err != nil || community == nil {
		t.Errorf("Not expected an error to get a community that should exists in the database: %s", err)
	}

	return community
}

func testCommunityAdapterInvalidGet(t *testing.T, adapter adapters.CommunityAdapter, id string) {
	_, err := adapter.Get(id)

	if err == nil {
		t.Errorf("Expected an error to get a community that shouldn't exists in the database: %s", err)
	}
}

func testCommunityAdapterName(t *testing.T, adapter adapters.CommunityAdapter, name string) *models.Community {
	community, err := adapter.Name(name)

	if err != nil {
		t.Errorf("Not expected an error to get a community by its name that should exists in the database: %s", err)
	}

	return community
}

func testCommunityAdapterInvalidName(t *testing.T, adapter adapters.CommunityAdapter, name string) {
	_, err := adapter.Name(name)

	if err == nil {
		t.Errorf("Expected an error to get a community by its name that shouldn't exists in the database: %s", err)
	}
}

func testCommunityAdapterSearch(t *testing.T, adapter adapters.CommunityAdapter, text string) {
	communities, err := adapter.Search(text)

	if err != nil {
		t.Errorf("Not expected an error to search communities: %s", err)
	}

	if len(communities) < 1 {
		t.Errorf("Not expected an empty array as the result of the search: %s", err)
	}
}

func testCommunityAdapterAll(t *testing.T, adapter adapters.CommunityAdapter) []models.Community {
	communities, err := adapter.All()

	if err != nil {
		t.Errorf("Not expected an error to get all the communities registered in the database: %s", err)
	}

	return communities
}

func testCommunityAdapterValidPublicCommunitiesByMultipleIds(t *testing.T, adapter adapters.CommunityAdapter, ids []string) []models.Community {
	communities, err := adapter.GetCommunitiesThroughIDS(ids)

	if err != nil {
		t.Errorf("Not expected an error to get public communities at the same time: %s", err)
	}

	if len(communities) < 1 {
		t.Errorf("Expected an array with multiple objects(models.Community), but got an empty array: %s", err)
	}

	return communities
}

func testCommunityAdapter(t *testing.T, adapter adapters.CommunityAdapter) {
	user, userErr := models.NewUser("choco", "choco@choco", []byte("choco"))

	if userErr != nil {
		t.Errorf("Not expected an error to create the user model to test the community adapter: %s", userErr)
	}

	community, commErr := models.NewCommunity("Choco", "Group to test things a.a", user.ID, false)

	if commErr != nil {
		t.Errorf("Not expected an error to create the community model to test the community adapter: %s", commErr)
	}

	testCommunityAdapterCreate(t, adapter, community)
	testCommunityAdapterGet(t, adapter, community.ID)
	testCommunityAdapterInvalidGet(t, adapter, "pdkawpodkad")
	testCommunityAdapterName(t, adapter, community.Name)
	testCommunityAdapterInvalidName(t, adapter, "dpkapkapapapapapapa")
	testCommunityAdapterSearch(t, adapter, community.Name)
	communities := testCommunityAdapterAll(t, adapter)

	// t.Errorf("%v\n", communities)

	testCommunityAdapterValidPublicCommunitiesByMultipleIds(t, adapter, []string{communities[0].ID})
}
