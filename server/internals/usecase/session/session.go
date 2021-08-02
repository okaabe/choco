package session

import (
	"choco/server/internals/adapters"
	"choco/server/internals/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type SessionUseCase struct {
	Issuer []byte
	Secret []byte

	UserAdapter adapters.UserAdapter
}

func (this *SessionUseCase) Register(username, email string, password []byte) (*models.User, string, error) {
	hash, hashErr := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if hashErr != nil {
		return nil, "", errors.New("Couldn't hash the password")
	}

	user, userErr := models.NewUser(username, email, hash)

	if userErr != nil {
		return nil, "", errors.New("Couldn't register the user")
	}

	adapterErr := this.UserAdapter.Add(user)

	if adapterErr != nil {
		return nil, "", errors.New("Couldn't register the user in the database")
	}

	token, jwtErr := Encode(this.Secret, string(this.Issuer), email)

	if jwtErr != nil {
		return nil, "", errors.New("Couldn't create the jwt.")
	}

	return user, token, nil
}

func (this *SessionUseCase) Authenticate(email string, password []byte) (*models.User, string, error) {
	user, userErr := this.UserAdapter.Email(email)

	if userErr != nil {
		return nil, "", errors.New("Couldn't find the user that has this email registered.")
	}

	isOk := bcrypt.CompareHashAndPassword(user.Password, password)

	if isOk != nil {
		return nil, "", errors.New("Incorrect password")
	}

	token, jwtErr := Encode(this.Secret, string(this.Issuer), email)

	if jwtErr != nil {
		return nil, "", errors.New("Couldn't create the jwt")
	}

	return user, token, nil
}

func (this *SessionUseCase) Rewoke(token string) (*models.User, error) {
	claim, jwtErr := IsExpiredAndDecode(token, this.Secret)

	if jwtErr != nil {
		return nil, jwtErr
	}

	user, userErr := this.UserAdapter.Email(claim.Email)

	if userErr != nil {
		return nil, errors.New("Couldn't find any user registered using this email")
	}

	return user, nil
}
