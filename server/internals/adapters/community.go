package adapters

import (
	"choco/server/internals/models"
	"fmt"

	"gorm.io/gorm"
)

type CommunityAdapter interface {
	Add(community *models.Community) error
	Get(id string) (*models.Community, error)
	Name(name string) (*models.Community, error)
	Search(text string) ([]models.Community, error)
}

type CommunityAdapterImpl struct {
	Adapter *gorm.DB
}

func (this *CommunityAdapterImpl) Add(community *models.Community) error {
	return this.Adapter.Create(community).Error
}

func (this *CommunityAdapterImpl) get(property string, value string) (*models.Community, error) {
	var community models.Community

	err := this.Adapter.First(&community, fmt.Sprintf("%s = ?", property), value).Error

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (this *CommunityAdapterImpl) Get(id string) (*models.Community, error) {
	return this.get("id", id)
}

func (this *CommunityAdapterImpl) Name(name string) (*models.Community, error) {
	return this.get("name", name)
}

func (this *CommunityAdapterImpl) Search(text string) ([]models.Community, error) {
	var communities []models.Community

	err := this.Adapter.Where("name LIKE ? OR description LIKE ?", "%" + text + "%", "%" + text + "%").Find(&communities).Error

	if err != nil {
		return nil, err
	}

	return communities, nil
}
