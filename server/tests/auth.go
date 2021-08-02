package tests

import (
	"choco/server/internals/models"
	"choco/server/internals/session"
	"testing"
)

func testAuthRegister(t *testing.T, auth *session.SessionUseCase, username, email string, password []byte) (*models.User, string) {
	user, token, err := auth.Register(username, email, password)

	if err != nil {
		t.Errorf("Not expected an error to register an user: %s", err)
	}

	return user, token
}

func testAuthRewoke(t *testing.T, auth *session.SessionUseCase, token string) *models.User {
	user, err := auth.Rewoke(token)

	if err != nil {
		t.Errorf("Not expected an error to rewoke a token: %s", err)
	}

	return user
}

func testAuthInvalidRewoke(t *testing.T, auth *session.SessionUseCase, token string) {
	_, err := auth.Rewoke(token)

	if err == nil {
		t.Errorf("Expected an error to rewoke an invalid token: %s", err)
	}
}

func testAuthValidAuthenticate(t *testing.T, auth *session.SessionUseCase, email string, password []byte) string {
	_, token, err := auth.Authenticate(email, password)

	if err != nil {
		t.Errorf("Not expected an error to authenticate with a valid email and password: %s", err)
	}

	return token
}

func testAuthInvalidAuthenticate(t *testing.T, auth *session.SessionUseCase, email string, password []byte) {
	_, _, err := auth.Authenticate(email, password)

	if err == nil {
		t.Errorf("Expected an error to authenticate with an incorrect email and password: %s", err)
	}
}

func testAuth(t *testing.T, auth *session.SessionUseCase) {
	var (
		username = "choco"
		email    = "choco@choco"
		password = []byte("choco")
	)

	_, token := testAuthRegister(t, auth, username, email, password)

	testAuthRewoke(t, auth, token)
	testAuthInvalidRewoke(t, auth, "a.a.a.a.a.aa.a")

	testAuthValidAuthenticate(t, auth, email, password)
	testAuthInvalidAuthenticate(t, auth, "aa..a.a.a@a.a.a.a", []byte("aoaoaoa"))
}
