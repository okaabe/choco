package inputs

type CreatePost struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func (this *CreatePost) Invalid() bool {
	return this.Text == "" || this.Title == ""
}
