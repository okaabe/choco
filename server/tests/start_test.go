package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/content"
	"choco/server/internals/session"
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
	testMemberAdapter(t, &adapters.MemberAdapterImpl{
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

	var sessionUseCase = &session.SessionUseCase{
		Issuer:      []byte("test"),
		Secret:      []byte("test"),
		UserAdapter: userAdapter,
	}

	testAuth(t, sessionUseCase)

	testContent(t, sessionUseCase, &content.Content{
		Session: sessionUseCase,

		MemberAdapter: &adapters.MemberAdapterImpl{
			Adapter: adapter,
		},

		CommunityAdapter: &adapters.CommunityAdapterImpl{
			Adapter: adapter,
		},

		PostAdapter: &adapters.PostAdapterImpl{
			Adapter: adapter,
		},
	})
}
