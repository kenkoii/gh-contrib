package usecases

import (
	"github.com/kenkoii/a-fis-gh-contrib/models"
	"github.com/kenkoii/a-fis-gh-contrib/repositories"
)

type UserUsecaseInterface interface {
	Save(user *models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Delete(username string) error
}

type UserUsecase struct {
	Repo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		Repo: repo,
	}
}

func (u *UserUsecase) Save(user *models.User) (*models.User, error) {
	return u.Repo.Save(user)
}

func (u *UserUsecase) GetByUsername(username string) (*models.User, error) {
	return u.Repo.GetByUsername(username)
}

func (u *UserUsecase) GetAll() ([]*models.User, error) {
	return u.Repo.GetAll()
}

func (u *UserUsecase) Delete(username string) error {
	return u.Repo.Delete(username)
}
