package adapter

import (
	"choco/server/internal/adapter/impl"
	"choco/server/internal/models"

	"gorm.io/gorm"
)

type UserAdapter interface {
	Add(*models.User) error

	GetByEmail(email string) (*models.User, error)
	GetByID(id string) (*models.User, error)

	DeleteByID(id string) error
	DeleteByEmail(email string) error

	UpdatePermissionByEmail(email string, permission int) error
	UpdatePermissionByID(id string, permission int) error
}

func NewUserAdapter(db *gorm.DB) UserAdapter {
	return &impl.UserAdapterImpl{
		DB: db,
	}
}