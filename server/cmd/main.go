package cmd

import (
	"choco/server/internals/http"
	"os"
)

func RunApp() {
	Environment()

	http.RunServer(os.Getenv("SERVER_ADDRESS"))
}
