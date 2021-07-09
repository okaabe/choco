package tests

import (
	"choco/server/internals/auth"
	"choco/server/internals/models"
	"testing"
)

func testAuthRegister(t *testing.T, auth *auth.Auth, email string, password []byte) (*models.User, string) {
	return nil, ""
}

func testAuthRewoke(t *testing.T, auth *auth.Auth, token string) *models.User {
	return nil
}

func testAuthInvalidRewoke(t *testing.T, auth *auth.Auth, token string) {
}

func testAuthValidAuthenticate(t *testing.T, auth *auth.Auth, email string, password []byte) string {
	return ""
}

func testAuthInvalidAuthenticate(t *testing.T, auth *auth.Auth, email string, password []byte) {
}

func testAuth(t *testing.T, auth *auth.Auth) {
}
