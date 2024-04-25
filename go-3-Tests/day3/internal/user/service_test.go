package users

import (
	"testing"

	. "github.com/luizlcezario/MelicampGoWeb/internal/domain"
	"github.com/luizlcezario/MelicampGoWeb/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func CreateService(t *testing.T) (*mocks.MockUserRepository, Service) {
	t.Helper()
	rep := new(mocks.MockUserRepository)
	service := UserService{rep}
	return rep, &service
}

func TestServiceUpdate(t *testing.T) {
	t.Run("try to Update The first user", func(t *testing.T) {
		rep, service := CreateService(t)
		expect := listExpect(t)
		userUpdt := User{Id: "", Name: "pokemon", SourName: "lima", Email: "pok@email", Age: 0, Heigth: 0.0}
		expectUser := User{Id: expect[0].Id, Name: userUpdt.Name, SourName: userUpdt.SourName, Email: userUpdt.Email, Age: userUpdt.Age, Heigth: userUpdt.Heigth, IsActive: userUpdt.IsActive, CreateadAt: userUpdt.CreateadAt}

		rep.On("FindById", expect[0].Id).Return(expect[0], nil)
		rep.On("Update", expect[0], userUpdt).Return(expectUser, nil)
		user, err := service.UpdateUser(expect[0].Id, "pokemon", "lima", "pok@email", 0, 0.0)

		assert.Nil(t, err)
		assert.Equal(t, expectUser, user)
		rep.AssertCalled(t, "FindById", expect[0])
		rep.AssertCalled(t, "Update", expect[0], userUpdt)
	})
}

// 	t.Run("testing service with user not change", func(t *testing.T) {
// 		mock := store.MockStore{
// 			WasCalledRead:  false,
// 			WasCalledWrite: false,
// 			Data:           expect,
// 			Rt:             nil,
// 		}
// 		repository := UserRepository{db: &mock}
// 		service := NewService(&repository)
// 		user, err := service.UpdateUser(expect[0].Id, "luiz", "lima", "luiz@gmail.com", 20, 1.7)
// 		user.CreateadAt = expect[0].CreateadAt

// 		assert.Nil(t, err)
// 		assert.Len(t, mock.Data, 1)
// 		assert.Equal(t, expect, mock.Data)
// 		assert.Equal(t, expect[0], user)
// 		assert.Equal(t, true, mock.WasCalledRead)
// 		assert.Equal(t, true, mock.WasCalledWrite)
// 	})

// 	t.Run("testing service with user change", func(t *testing.T) {
// 		mock := store.MockStore{
// 			WasCalledRead:  false,
// 			WasCalledWrite: false,
// 			Data:           expect,
// 			Rt:             nil,
// 		}
// 		repository := UserRepository{db: &mock}
// 		service := NewService(&repository)
// 		expectUser := User{expect[0].Id, "vinicinho", "pinto", "pinto@gmail.com", 21, 1.8, true, expect[0].CreateadAt}

// 		user, err := service.UpdateUser(expect[0].Id, "vinicinho", "pinto", "pinto@gmail.com", 21, 1.8)
// 		user.CreateadAt = expect[0].CreateadAt

// 		assert.Nil(t, err)
// 		assert.Len(t, mock.Data, 1)
// 		assert.Equal(t, expect, mock.Data)
// 		assert.Equal(t, expectUser, user)
// 		assert.Equal(t, true, mock.WasCalledRead)
// 		assert.Equal(t, true, mock.WasCalledWrite)
// 	})

// 	t.Run("testing service reciven error", func(t *testing.T) {
// 		mock := store.MockStore{
// 			WasCalledRead:  false,
// 			WasCalledWrite: false,
// 			Data:           nil,
// 			Rt:             errors.New("error db"),
// 		}
// 		repository := UserRepository{db: &mock}
// 		service := NewService(&repository)

// 		user, err := service.UpdateUser("asdasd", "vinicinho", "pinto", "pinto@gmail.com", 21, 1.8)

// 		assert.NotNil(t, err)
// 		assert.Error(t, err)
// 		assert.Equal(t, User{}, user)
// 		assert.Equal(t, true, mock.WasCalledRead)
// 		assert.Equal(t, false, mock.WasCalledWrite)
// 	})
// }

// func TestServiceDelete(t *testing.T) {
// 	expect := []User{
// 		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
// 	}

// 	t.Run("Deleting an existing User", func(t *testing.T) {
// 		mock := store.MockStore{
// 			WasCalledRead:  false,
// 			WasCalledWrite: false,
// 			Data:           expect,
// 			Rt:             nil,
// 		}
// 		repository := UserRepository{db: &mock}
// 		service := NewService(&repository)

// 		err := service.Delete(expect[0].Id)

// 		assert.Nil(t, err)
// 		assert.Equal(t, true, mock.WasCalledRead)
// 		assert.Equal(t, true, mock.WasCalledWrite)
// 	})

// 	t.Run("Deleting an existing User", func(t *testing.T) {
// 		mock := store.MockStore{
// 			WasCalledRead:  false,
// 			WasCalledWrite: false,
// 			Data:           expect,
// 			Rt:             errors.New("error db"),
// 		}
// 		repository := UserRepository{db: &mock}
// 		service := NewService(&repository)

// 		err := service.Delete(expect[0].Id)

// 		assert.NotNil(t, err)
// 		assert.Error(t, err)
// 		assert.Equal(t, true, mock.WasCalledRead)
// 		assert.Equal(t, false, mock.WasCalledWrite)
// 	})
// }
