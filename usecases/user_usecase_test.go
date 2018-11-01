package usecases

import (
	"testing"

	"github.com/kenkoii/a-fis-gh-contrib/models"
	"github.com/kenkoii/a-fis-gh-contrib/repositories"
)

func TestSave(t *testing.T) {
	r := repositories.NewInMemoryUserRepository()
	u := NewUserUsecase(r)

	user := &models.User{
		Username: "kenkoii",
		ImageUrl: "https://avatars0.githubusercontent.com/u/11359825?s=460&v=4",
	}
	_, err := u.Save(user)
	if err != nil {
		t.Errorf("Error saving, message: %s", err.Error())
	}
}
