package impl

import (
	"choco/server/internal/models"

	"gorm.io/gorm"
)

type UserAdapterImpl struct {
	DB *gorm.DB
}

func (adapter *UserAdapterImpl) first(query ...interface{}) (*models.User, error) {
	var user models.User

	err := adapter.DB.First(&user, query...).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (adapter *UserAdapterImpl) delete(query ...interface{}) error {
	return adapter.DB.Delete(&models.User{}, query...).Error
}

func (adapter *UserAdapterImpl) update_permission(permission int, query ...interface{}) error {
	return adapter.DB.Model(&models.User{}).Where(query).Update("permission", permission).Error
}

func (adapter *UserAdapterImpl) Add(user *models.User) error {
	return adapter.DB.Save(user).Error
}

func (adapter *UserAdapterImpl) GetByEmail(email string) (*models.User, error) {
	return adapter.first("email = ?", email)
}

func (adapter *UserAdapterImpl) GetByID(id string) (*models.User, error) {
	return adapter.first("id = ?", id)
}

func (adapter *UserAdapterImpl) DeleteByEmail(email string) error {
	return adapter.delete("email = ?", email)
}

func (adapter *UserAdapterImpl) DeleteByID(id string) error {
	return adapter.delete("id = ?", id)
}

func (adapter *UserAdapterImpl) UpdatePermissionByEmail(email string, permission int) error {
	return adapter.update_permission(permission, "email = ?", email)
}

func (adapter *UserAdapterImpl) UpdatePermissionByID(id string, permission int) error {
	return adapter.update_permission(permission, "id = ?", id)
}
