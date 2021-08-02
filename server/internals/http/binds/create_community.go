package binds

type CreateCommunity struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Nsfw        bool   `json:"nsfw"`
}

func (this *CreateCommunity) Invalid() bool {
	return this.Name == "" || this.Description == ""
}
