package adapters

import "choco/server/internals/models"

type PostAdapter interface {
	Add(post *models.Post) error
	Get(id string) (*models.Post, error)
	Community(community_id string) ([]models.Community, error)
	Delete(id string)
}