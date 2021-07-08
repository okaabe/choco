package models

type Community struct {
	Base
	Name        string `json:"name" gorm:"unique;"`
	CreatorID   string `json:"creator_id"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
	Nsfw        bool   `json:"nsfw"`
}

func NewCommunity(name, description, creator_id string, private, nsfw bool) (*Community, error) {
	community := Community{
		Name:        name,
		Description: description,
		CreatorID:   creator_id,
		Private:     private,
		Nsfw:        nsfw,
	}

	community.NewBase()

	return &community, nil
}
