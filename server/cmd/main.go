package cmd

import (
	"choco/server/internals/adapters"
	"choco/server/internals/auth"
	"choco/server/internals/content"
	"choco/server/internals/http"
	"os"
)

func RunApp() {
	Environment()

	var (
		adapter = adapters.ConnectTestDB()

		auth = &auth.Auth{
			Issuer: []byte(os.Getenv("JWT_ISSUER")),
			Secret: []byte(os.Getenv("JWT_SECRET")),
			UserAdapter: &adapters.UserAdapterImpl{
				Adapter: adapter,
			},
		}

		content = &content.Content{
			Auth:             auth,
			CommunityAdapter: &adapters.CommunityAdapterImpl{Adapter: adapter},
			PostAdapter:      &adapters.PostAdapterImpl{Adapter: adapter},
		}
	)

	http.RunServer(os.Getenv("SERVER_ADDRESS"), auth, content)
}
