package models

type User struct {
	Base
	Username   string `json:"username"`
	Email      string `json:"email" gorm:"unique;"`
	Password   []byte `json:"password"`
}

func NewUser(username, email string, password []byte) (*User, error) {
	user := User{
		Username:   username,
		Password:   password,
		Email:      email,
	}

	user.NewBase()

	return &user, nil
}
