package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
	// "choco/server/internals/auth"
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
	testPostAdapter(t, &adapters.PostAdapterImpl{
		Adapter: adapter,
	})
}

func TestAuthJwt(t *testing.T) {
	key := []byte("chocotestjwtsla")
	issuer := "oioioioioi"

	token := testAuthJwtEncode(t, key, issuer)

	testAuthJwtValidDecode(t, key, token)
	testAuthJwtInvalidDecode(t, key, "opdakwpdokskdzlxdçzklçldka.doawkdpoawkd.dwkaopdawd")
}

func TestFeatures(t *testing.T) {
	adapter := adapters.ConnectTestDB()
	var userAdapter = &adapters.UserAdapterImpl{
		Adapter: adapter,
	}

	testAuth(t, &auth.Auth{
		Issuer: []byte("test"),
		Secret: []byte("test"),
		UserAdapter: userAdapter,
	})
}