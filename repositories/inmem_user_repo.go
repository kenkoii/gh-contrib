package repositories

import (
	"fmt"

	"github.com/kenkoii/a-fis-gh-contrib/models"
)

type InMemoryUserRepository struct {
	data map[string]*models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]*models.User),
	}
}

func (repo *InMemoryUserRepository) Save(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, fmt.Errorf("Object is nil, cannot be inserted")
	}
	repo.data[user.Username] = user
	return user, nil
}

func (repo *InMemoryUserRepository) GetByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, fmt.Errorf("Username must not be empty")
	}
	if repo.data[username] == nil {
		return nil, fmt.Errorf("Username not found")
	}

	return repo.data[username], nil
}

func (repo *InMemoryUserRepository) GetAll() ([]*models.User, error) {
	var list []*models.User
	for _, j := range repo.data {
		list = append(list, j)
	}
	return list, nil
}

func (repo *InMemoryUserRepository) Delete(username string) error {
	if username == "" {
		return fmt.Errorf("Username must not be empty")
	}

	if repo.data[username] == nil {
		return fmt.Errorf("Username not found")
	}

	delete(repo.data, username)
	return nil
}
