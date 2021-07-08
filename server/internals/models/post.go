package models

type Post struct {
	Base
	Title       string `json:"title"`        // Title of the post
	Text        string `json:"text"`         // the text of the post or if you'd  rather the body of the post
	AuthorID    string `json:"author_id"`    // That's the id of the user who created the post
	CommunityID string `json:"community_id"` //That's the id of the community where is the post, for example (unixporn and etc)
	Private     bool   `json:"private"`      // That's a property that will define if the post is private or not
	Nsfw        bool   `json:"nsfw"`         // The same of the previous but about the nsfw or in other words if the post is +18
}

func NewPost(title, text, author_id, community_id string, private, nsfw bool) (*Post, error) {
	post := Post{
		Title:    title,
		Text:     text,
		AuthorID: author_id,
		CommunityID: community_id,
		Private:  private,
		Nsfw:     nsfw,
	}

	post.NewBase()

	return &post, nil
}
