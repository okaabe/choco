package inputs

type CreateCommunity struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Nsfw        bool   `json:"nsfw"`
	Private     bool   `json:"private"`
}

func (this *CreateCommunity) Invalid() bool {
	return this.Name == "" || this.Description == ""
}
