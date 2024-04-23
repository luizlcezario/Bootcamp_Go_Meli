package users

import (
	"errors"
	"testing"

	"github.com/luizlcezario/MelicampGoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceUpdate(t *testing.T) {
	expect := []User{
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
	}

	t.Run("testing service with user not change", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           expect,
			Rt:             nil,
		}
		repository := UserRepository{db: &mock}
		service := NewService(&repository)
		user, err := service.UpdateUser(expect[0].Id, "luiz", "lima", "luiz@gmail.com", 20, 1.7)
		user.CreateadAt = expect[0].CreateadAt

		assert.Nil(t, err)
		assert.Len(t, mock.Data, 1)
		assert.Equal(t, expect, mock.Data)
		assert.Equal(t, expect[0], user)
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, true, mock.WasCalledWrite)
	})

	t.Run("testing service with user change", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           expect,
			Rt:             nil,
		}
		repository := UserRepository{db: &mock}
		service := NewService(&repository)
		expectUser := User{expect[0].Id, "vinicinho", "pinto", "pinto@gmail.com", 21, 1.8, true, expect[0].CreateadAt}

		user, err := service.UpdateUser(expect[0].Id, "vinicinho", "pinto", "pinto@gmail.com", 21, 1.8)
		user.CreateadAt = expect[0].CreateadAt

		assert.Nil(t, err)
		assert.Len(t, mock.Data, 1)
		assert.Equal(t, expect, mock.Data)
		assert.Equal(t, expectUser, user)
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, true, mock.WasCalledWrite)
	})

	t.Run("testing service reciven error", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           nil,
			Rt:             errors.New("error db"),
		}
		repository := UserRepository{db: &mock}
		service := NewService(&repository)

		user, err := service.UpdateUser("asdasd", "vinicinho", "pinto", "pinto@gmail.com", 21, 1.8)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, User{}, user)
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, false, mock.WasCalledWrite)
	})
}

func TestServiceDelete(t *testing.T) {
	expect := []User{
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
	}

	t.Run("Deleting an existing User", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           expect,
			Rt:             nil,
		}
		repository := UserRepository{db: &mock}
		service := NewService(&repository)

		err := service.Delete(expect[0].Id)

		assert.Nil(t, err)
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, true, mock.WasCalledWrite)
	})

	t.Run("Deleting an existing User", func(t *testing.T) {
		mock := store.MockStore{
			WasCalledRead:  false,
			WasCalledWrite: false,
			Data:           expect,
			Rt:             errors.New("error db"),
		}
		repository := UserRepository{db: &mock}
		service := NewService(&repository)

		err := service.Delete(expect[0].Id)

		assert.NotNil(t, err)
		assert.Error(t, err)
		assert.Equal(t, true, mock.WasCalledRead)
		assert.Equal(t, false, mock.WasCalledWrite)
	})
}
