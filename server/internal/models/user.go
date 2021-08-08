package models

const (
	Normal            = 1
	Moderator     int = 6
	Administrator int = 9
)

type User struct {
	Base
	Username   string `json:"username"`
	Email      string `json:"email" gorm:"unique; type:email;"`
	Password   []byte `json:"password"`
	Biograph   string `json:"biograph"`
	Permission int    `json:"permission"`
}

func (user *User) Invalid() bool {
	return user.Permission != Normal && user.Permission != Moderator && user.Permission != Administrator
}

func NewUser(username, email, biograph string, password []byte, permission int) (*User, error) {
	user := User{
		Username:   username,
		Email:      email,
		Password:   password,
		Biograph:   biograph,
		Permission: permission,
	}

	err := user.id()

	if err != nil {
		return nil, err
	}

	return &user, nil
}
