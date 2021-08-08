package adapter

import (
	"choco/server/internal/models"
)

type UserAdapter interface {
	Add(*models.User) *models.User

	GetByEmail(email string) (*models.User, error)
	GetByID(id string) (*models.User, error)

	DeleteByID(id string) error
	DeleteByEmail(email string) error

	UpdatePermissionByEmail(email string, permission int) error
	UpdatePermissionByID(id string, permission int) error
}
