package adapters

import (
	"choco/server/internals/models"
	"fmt"

	"gorm.io/gorm"
)

type UserAdapter interface {
	Add(*models.User) error
	Get(id string) (*models.User, error)
	Email(email string) (*models.User, error)
	Delete(id string) error
}

type UserAdapterImpl struct {
	Adapter *gorm.DB
}

func (this *UserAdapterImpl) Add(user *models.User) error {
	return this.Adapter.Create(user).Error
}

func (this *UserAdapterImpl) get(prop string, value interface{}) (*models.User, error) {
	var user models.User

	err := this.Adapter.First(&user, fmt.Sprintf("%s = ?", prop), value).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (this *UserAdapterImpl) Get(id string) (*models.User, error) {
	return this.get("id", id)
}

func (this *UserAdapterImpl) Email(email string) (*models.User, error) {
	return this.get("email", email)
}

func (this *UserAdapterImpl) Delete(id string) error {
	return this.Adapter.Delete(&models.User{}, "id = ?", id).Error
}
