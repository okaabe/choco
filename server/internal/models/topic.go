package models

type Topic struct {
	Base
	Name   string `json:"name" gorm:"unique"`
	About  string `json:"about"`
	UserID string `json:"user_id"`
}

func NewTopic(name, about, user_id string, nsfw bool) (*Topic, error) {
	topic := Topic{
		Name:   name,
		About:  about,
		UserID: user_id,
	}

	err := topic.id()

	if err != nil {
		return nil, err
	}

	return &topic, nil
}
