package usecases

import (
	"testing"

	"github.com/kenkoii/a-fis-gh-contrib/models"
	"github.com/kenkoii/a-fis-gh-contrib/repositories"
)

func TestUsecase(t *testing.T) {
	r := repositories.NewInMemoryUserRepository()
	u := NewUserUsecase(r)

	user := &models.User{
		Username: "kenkoii",
		ImageUrl: "https://avatars0.githubusercontent.com/u/11359825?s=460&v=4",
	}

	t.Run("Save", func(t *testing.T) {
		_, err := u.Save(user)
		if err != nil {
			t.Errorf("Error saving, message: %s", err.Error())
		}
	})

	t.Run("GetByUsername", func(t *testing.T) {
		want := "kenkoii"
		got, err := u.GetByUsername(want)
		if err != nil {
			t.Errorf("No user fetched error: %s", err.Error())
		}

		if got.Username != want {
			t.Errorf("Wrong user fetched: %s", err.Error())
		}
	})
}
