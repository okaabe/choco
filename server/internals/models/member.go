package models

type Member struct {
	Base
	UserID      string `json:"user_id"`
	CommunityID string `json:"community_id"`
}

func NewMember(userId, communityId string) (*Member, error) {
	member := Member{
		UserID:      userId,
		CommunityID: communityId,
	}

	member.NewBase()

	return &member, nil
}
