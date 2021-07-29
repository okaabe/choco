package adapters

import (
	"choco/server/internals/models"

	"gorm.io/gorm"
)

type PostAdapter interface {
	Add(post *models.Post) error
	Get(id string) (*models.Post, error)
	Community(community_id string) ([]models.Post, error)
	Search(text string) ([]models.Post, error)
}

type PostAdapterImpl struct {
	Adapter *gorm.DB
}

func (this *PostAdapterImpl) Add(post *models.Post) error {
	return this.Adapter.Create(post).Error
}

func (this *PostAdapterImpl) get(query ...interface{}) (*models.Post, error) {
	var post models.Post

	err := this.Adapter.First(&post, query...).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (this *PostAdapterImpl) find(query ...interface{}) ([]models.Post, error) {
	var posts []models.Post

	err := this.Adapter.Find(&posts, query...).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (this *PostAdapterImpl) Get(id string) (*models.Post, error) {
	return this.get("id = ?", id)
}

func (this *PostAdapterImpl) Community(community_id string) ([]models.Post, error) {
	return this.find("community_id = ?", community_id)
}

func (this *PostAdapterImpl) Search(text string) ([]models.Post, error) {
	return this.find("title LIKE ? OR text LIKE ?", "%"+text+"%", "%"+text+"%")
}
