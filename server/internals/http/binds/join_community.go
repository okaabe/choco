package binds

type JoinCommunity struct {
	CommunityName string `json:"community_name"`
}

func (this *JoinCommunity) Invalid() bool {
	return this.CommunityName == ""
}
