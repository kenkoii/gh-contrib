package repositories

import "github.com/kenkoii/gh-contrib/models"

type UserRepository interface {
	Save(user *models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(username string) error
}
