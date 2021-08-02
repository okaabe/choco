package binds

type CreatePost struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Nsfw  bool   `json:"nsfw"`
}

func (this *CreatePost) Invalid() bool {
	return this.Text == "" || this.Title == ""
}
