package tests

import (
	"choco/server/internals/auth"
	"choco/server/internals/content"
	"choco/server/internals/models"
	"testing"
)

func testContentValidCreateCommunity(t *testing.T, content *content.Content, validToken string) *models.Community {
	community, communityErr := content.CreateCommunity("Choco tests", "A simple community to work as...", validToken, false, false)

	if communityErr != nil {
		t.Errorf("Not expected an error to create a community: %s\n", communityErr)
	}

	return community
}

func testContentInvalidCreateCommunity(t *testing.T, content *content.Content, invalidToken string) {

}

func testContent(t *testing.T, auth *auth.Auth, content *content.Content) {
	_, token, registerErr := auth.Register("okaabe", "okaabe@okaabe", []byte("okaabe"))

	if registerErr != nil {
		t.Errorf("Not expected an error to register an user to test the content features")
	}

	invalidToken := "apkwpkad.kdpwoakdpawkdpakdw.dkwpoadkwd"

	testContentValidCreateCommunity(t, content, token)
	testContentInvalidCreateCommunity(t, content, invalidToken)
}
