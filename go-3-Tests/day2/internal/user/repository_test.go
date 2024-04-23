package users

import (
	"testing"
	"time"

	"github.com/luizlcezario/MelicampGoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Testing GetAll with [] without error of repository", func(t *testing.T) {
		var expect []User
		stub := store.StubStore{}
		repository := UserRepository{db: &stub}
		user, err := repository.GetAll()

		assert.Equal(t, expect, user)
		assert.Nil(t, err)
	})

}

func TestRepositoryUpdate(t *testing.T) {
	expect := []User{
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
	}
	t.Run("Testing Update nothing changes", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           &expect,
			Rt:             nil,
		}

		repository := UserRepository{db: &mock}

		repository.Update(expect[0], expect[0])

		assert.Len(t, mock.Data, 1)
		assert.Equal(t, expect, mock.Data)
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, true, mock.WasCalledWrite)
	})

	t.Run("Testing Update with changes", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
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
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, true, mock.WasCalledWrite)
	})
}

func TestRespositoryFindById(t *testing.T) {
	expect := []User{
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
	}
	t.Run("find first User", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           &expect,
			Rt:             nil,
		}
		repository := UserRepository{db: &mock}
		result := expect[1]

		user, err := repository.FindById(result.Id)
		result.CreateadAt = time.Now()
		result.CreateadAt = user.CreateadAt

		assert.Nil(t, err)
		assert.Equal(t, result, user)
	})

}
