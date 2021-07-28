package tests

import (
	"choco/server/internals/adapters"
	"choco/server/internals/models"
	"testing"
)

func testUserAdapterCreate(t *testing.T, adapter adapters.UserAdapter, object *models.User) *models.User {
	err := adapter.Add(object)

	if err != nil {
		t.Fatalf("Not expected an error to create an object using user adapter: %s", err)
	}

	return object
}

func testUserAdapterGet(t *testing.T, adapter adapters.UserAdapter, id string) *models.User {
	user, err := adapter.Get(id)

	if err != nil {
		t.Errorf("Not expected an error to get an user that should exists in the database: %s", err)
	}

	return user
}

func testUserAdapterInvalidGet(t *testing.T, adapter adapters.UserAdapter, id string) {
	_, err := adapter.Get(id)

	if err == nil {
		t.Errorf("Expected an error to get an user that shouldn't exists in the database: %s", err)
	}
}

func testUserAdapterEmailGet(t *testing.T, adapter adapters.UserAdapter, email string) *models.User {
	user, err := adapter.Email(email)

	if err != nil {
		t.Errorf("Not expected an error to get an user by its email, it should exists in the database: %s", err)
	}

	return user
}

func testUserAdapterInvalidEmailGet(t *testing.T, adapter adapters.UserAdapter, email string) {
	_, err := adapter.Email(email)

	if err == nil {
		t.Errorf("Expected an error to get an user by its email, it shouldn't exists in the dataabase: %s", err)
	}
}

func testUserAdapter(t *testing.T, adapter adapters.UserAdapter) {
	user, userErr := models.NewUser("choco", "choco@choco", []byte("choco"))

	if userErr != nil {
		t.Errorf("Not expected an error to create the user model: %s", userErr)
	}

	testUserAdapterCreate(t, adapter, user)
	testUserAdapterGet(t, adapter, user.ID)
	testUserAdapterInvalidGet(t, adapter, "podkawpdkapwdkawpd")
	testUserAdapterEmailGet(t, adapter, user.Email)
	testUserAdapterInvalidEmailGet(t, adapter, "pdkawpodkapwdk@pkdpawkdpaowkd.com")
}
