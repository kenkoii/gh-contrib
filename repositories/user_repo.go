package repositories

import "github.com/kenkoii/a-fis-gh-contrib/models"

type UserRepository interface {
	Save(user *models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Delete(username string) error
}
