package models

import "errors"

const (
	ROOT_PERMISSION uint = 999
	USER_PERMISSION uint = 10
)

type User struct {
	Base
	Username   string `json:"username"`
	Email      string `json:"email" gorm:"unique;"`
	Password   []byte `json:"password"`
	Permission uint   `json:"permission"`
}

func (this *User) validate() error {
	if this.Permission != ROOT_PERMISSION && this.Permission != USER_PERMISSION {
		return errors.New("Invalid permission")
	}

	return nil
}

func NewUser(username, email string, password []byte, permission uint) (*User, error) {
	user := User{
		Username:   username,
		Password:   password,
		Email:      email,
		Permission: permission,
	}

	err := user.validate()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
