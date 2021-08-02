package models

type Community struct {
	Base
	Name        string `json:"name" gorm:"unique;"`
	CreatorID   string `json:"creator_id"`
	Description string `json:"description"`
	Nsfw        bool   `json:"nsfw"`
}

func NewCommunity(name, description, creator_id string, nsfw bool) (*Community, error) {
	community := Community{
		Name:        name,
		Description: description,
		CreatorID:   creator_id,
		Nsfw:        nsfw,
	}

	community.NewBase()

	return &community, nil
}
