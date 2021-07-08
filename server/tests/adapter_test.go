package tests

import (
	"testing"
	"choco/server/internals/adapters"
)

func TestAdapters(t *testing.T) {
	var adapter = adapters.ConnectTestDB()

	testUserAdapter(t, &adapters.UserAdapterImpl{
		Adapter: adapter,
	})
}
