package tests

import (
	"choco/server/internals/data"
	"choco/server/internals/models"
	"testing"
)

func testInvalidUserModel(t *testing.T) *models.User {
	user, err := models.NewUser("choco", "choco@choco", []byte("choco"), 5444)

	if err != nil {
		t.Fatalf("Expected an error to create an invalid user model, but got nil")
	}

	return user
}

func testValidUserModel(t *testing.T) *models.User {
	user, err := models.NewUser("choco", "choco@choco", []byte("choco"), data.ADMIN_PERMISSION)

	if err != nil {
		t.Fatalf("Not expected an error to create a valid user, but got %s", err)
	}

	return user
}

func TestUserModel(t *testing.T) {
	testInvalidUserModel(t)
	testValidUserModel(t)
}
