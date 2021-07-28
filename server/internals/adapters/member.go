package adapters

import (
	"choco/server/internals/models"

	"gorm.io/gorm"
)

type MemberAdapter interface {
	Add(member *models.Member) error
	ID(id string) (*models.Member, error)
	UserID(id string) ([]models.Member, error)
	CommunityID(communityId string) ([]models.Member, error)
	MemberInTheCommunity(communityId, userId string) (*models.Member, error)
}

type MemberAdapterImpl struct {
	Adapter *gorm.DB
}

func (this *MemberAdapterImpl) first(query ...interface{}) (*models.Member, error) {
	var member models.Member

	err := this.Adapter.First(&member, query...).Error

	if err != nil {
		return nil, err
	}

	return &member, nil
}

func (this *MemberAdapterImpl) Add(member *models.Member) error {
	return this.Adapter.Create(member).Error
}

func (this *MemberAdapterImpl) ID(id string) (*models.Member, error) {
	return this.first("id = ?", id)
}

func (this *MemberAdapterImpl) MemberInTheCommunity(communityId, userId string) (*models.Member, error) {
	return nil, nil
}

func (this *MemberAdapterImpl) CommunityID(communityId string) ([]models.Member, error) {
	return nil, nil
}

func (this *MemberAdapterImpl) UserID(id string) ([]models.Member, error) {
	return nil, nil
}
