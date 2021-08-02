package binds

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this *SignUp) Invalid() bool {
	return this.Username == "" || this.Email == "" || this.Password == ""
}
