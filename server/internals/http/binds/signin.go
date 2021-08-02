package binds

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this *SignIn) Invalid() bool {
	return this.Email == "" || this.Password == ""
}
