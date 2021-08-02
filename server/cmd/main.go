package cmd

import (
	"choco/server/internals/adapters"
	"choco/server/internals/http"
	"choco/server/internals/usecase/content"
	"choco/server/internals/usecase/session"
	"os"
)

func RunApp() {
	Environment()

	var (
		adapter = adapters.ConnectTestDB()

		sessionUseCase = &session.SessionUseCase{
			Issuer: []byte(os.Getenv("JWT_ISSUER")),
			Secret: []byte(os.Getenv("JWT_SECRET")),
			UserAdapter: &adapters.UserAdapterImpl{
				Adapter: adapter,
			},
		}

		content = &content.Content{
			Session:          sessionUseCase,
			CommunityAdapter: &adapters.CommunityAdapterImpl{Adapter: adapter},
			PostAdapter:      &adapters.PostAdapterImpl{Adapter: adapter},
			MemberAdapter:    &adapters.MemberAdapterImpl{Adapter: adapter},
		}
	)

	http.RunServer(os.Getenv("SERVER_ADDRESS"), sessionUseCase, content)
}
