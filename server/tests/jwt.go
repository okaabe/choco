package tests

import (
	"choco/server/internals/usecase/session"
	"testing"
)

func testAuthJwtEncode(t *testing.T, key []byte, issuer string) string {
	token, err := session.Encode(key, issuer, "choco@choco")

	if err != nil {
		t.Errorf("Not expected an error to create an user token: %s", err)
	}

	return token
}

func testAuthJwtValidDecode(t *testing.T, key []byte, token string) *session.JwtClaim {
	claim, err := session.IsExpiredAndDecode(token, key)

	if err != nil {
		t.Errorf("Not expected an error to decode a valid jwt")
	}

	return claim
}

func testAuthJwtInvalidDecode(t *testing.T, key []byte, token string) *session.JwtClaim {
	claim, err := session.IsExpiredAndDecode(token, key)

	if err == nil {
		t.Errorf("Expected an error to decode an invalid jwt")
	}

	return claim
}
