package repositories

import (
	"testing"

	"github.com/kenkoii/a-fis-gh-contrib/models"
)

func TestInMemRepo(t *testing.T) {
	r := NewInMemoryUserRepository()

	user := &models.User{
		Username: "kenkoii",
		ImageUrl: "https://avatars0.githubusercontent.com/u/11359825?s=460&v=4",
	}

	t.Run("Save", func(t *testing.T) {
		_, err := r.Save(user)
		if err != nil {
			t.Errorf("Error saving, message: %s", err.Error())
		}
	})

	t.Run("GetByUsername", func(t *testing.T) {
		want := "kenkoii"
		got, err := r.GetByUsername(want)
		if err != nil {
			t.Errorf("No user fetched error: %s", err.Error())
		}

		if got.Username != want {
			t.Errorf("Wrong user fetched: %s", err.Error())
		}
	})

	t.Run("GetAll and Delete", func(t *testing.T) {
		want := 1
		got, err := r.GetAll()
		if err != nil {
			t.Errorf("Get all error: %s", err.Error())
		}

		if len(got) != want {
			t.Errorf("Wrong length: %s", err.Error())
		}

		want = 0
		err = r.Delete("kenkoii")
		if err != nil {
			t.Errorf("Delete error %s", err.Error())
		}

		got, err = r.GetAll()
		if err != nil {
			t.Errorf("Get All error %s", err.Error())
		}
		if len(got) != want {
			t.Errorf("Length should now be 0 %s", err.Error())
		}
	})
}
