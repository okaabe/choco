package tests

import (
	"choco/server/internals/adapters"
	"testing"
)

func TestAdapters(t *testing.T) {
	var adapter = adapters.ConnectTestDB()

	testUserAdapter(t, &adapters.UserAdapterImpl{
		Adapter: adapter,
	})
	testCommunityAdapter(t, &adapters.CommunityAdapterImpl{
		Adapter: adapter,
	})
}
