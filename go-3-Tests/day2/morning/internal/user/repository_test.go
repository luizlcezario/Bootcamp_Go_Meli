package users

import (
	"testing"

	"github.com/luizlcezario/MelicampGoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("Testing GetAll with [] without error of repository", func(t *testing.T) {
		var expect []User
		stub := store.StubStore{}
		repository := UserRepository{db: &stub}
		user, err := repository.GetAll()

		assert.Equal(t, expect, user)
		assert.Nil(t, err)
	})

}

func TestUpdate(t *testing.T) {
	expect := []User{
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
	}
	t.Run("Testing Update nothing changes", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  0,
			WasCalledWrite: 0,
			Data:           &expect,
			Rt:             nil,
		}

		repository := UserRepository{db: &mock}

		repository.Update(expect[0], expect[0])

		assert.Len(t, mock.Data, 1)
		assert.Equal(t, expect, mock.Data)
		assert.Equal(t, mock.WasCalledRead, 1)
		assert.Equal(t, mock.WasCalledWrite, 2)
	})

	t.Run("Testing Update with changes", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  0,
			WasCalledWrite: 0,
			Data:           &expect,
			Rt:             nil,
		}
		newUser := NewUser("vinicinho", "lima", "luiz@gmail.com", 20, 1.7, true)
		newUser.Id = ""
		repository := UserRepository{db: &mock}

		user := repository.Update(expect[0], *newUser)
		var result []User = []User{
			user,
		}

		assert.Len(t, mock.Data, 1)
		assert.Equal(t, result, mock.Data)
		assert.Equal(t, mock.WasCalledRead, 1)
		assert.Equal(t, mock.WasCalledWrite, 2)
	})

}
