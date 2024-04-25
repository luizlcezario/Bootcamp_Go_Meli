package users

import (
	"errors"
	"slices"
	"testing"

	. "github.com/luizlcezario/MelicampGoWeb/internal/domain"
	"github.com/luizlcezario/MelicampGoWeb/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func CreateRepository(t *testing.T) (*mocks.MockStore, Repository) {
	t.Helper()
	db := new(mocks.MockStore)
	repository := UserRepository{db: db}
	return db, &repository
}

func listExpect(t *testing.T) []User {
	t.Helper()
	return []User{
		*NewUser("luiz", "lima", "luiz@gmail.com", 20, 1.7, true),
		*NewUser("Milne", "Carecho", "mi@gmail.com", 23, 1.5, true),
		*NewUser("Luan", "Silva", "lu@gmail.com", 19, 1.6, true),
	}
}

func TestRepositoryGetAll(t *testing.T) {
	t.Run("Testing GetAll with [] without error of repository", func(t *testing.T) {
		var expect []User
		db, repository := CreateRepository(t)
		db.On("Read").Return(expect, nil).Once()
		user, err := repository.GetAll()

		assert.Equal(t, expect, user)
		assert.Len(t, user, 0)
		assert.Nil(t, err)
		db.AssertCalled(t, "Read")
	})

	t.Run("Testing GetAll with some array valid", func(t *testing.T) {
		expect := listExpect(t)
		db, repository := CreateRepository(t)
		db.On("Read").Return(expect, nil).Once()
		user, err := repository.GetAll()

		assert.Equal(t, expect, user)
		assert.Len(t, user, len(expect))
		assert.Nil(t, err)
	})

	t.Run("Testing error from db", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var expect []User
		db.On("Read").Return(expect, errors.New("tesarera	")).Once()
		user, err := repository.GetAll()

		db.AssertCalled(t, "Read")
		assert.NotNil(t, err)
		assert.Equal(t, expect, user)
	})

}

func TestRepositoryStore(t *testing.T) {
	t.Run("creating new user simple", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var nUser User = *NewUser("luiz", "lima", "luiz@gmail", 20, 1.7, true)
		var expect []User

		db.On("Read").Return(expect, nil).Once()
		expect = append(expect, nUser)
		db.On("Write", expect).Return(nil).Once()
		user, err := repository.Store(nUser)

		assert.Nil(t, err)
		assert.Equal(t, nUser, user)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", expect)
	})

	t.Run("creating new with something in the store", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var nUser User = *NewUser("luiz", "lima", "luiz@gmail", 20, 1.7, true)
		expect := listExpect(t)

		db.On("Read").Return(expect, nil).Once()
		expect = append(expect, nUser)
		db.On("Write", expect).Return(nil).Once()
		user, err := repository.Store(nUser)

		assert.Nil(t, err)
		assert.Equal(t, nUser, user)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", expect)
	})
	t.Run("creating new with error in the read", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var nUser User = *NewUser("luiz", "lima", "luiz@gmail", 20, 1.7, true)
		expect := listExpect(t)

		db.On("Read").Return(expect, errors.New("")).Once()
		_, err := repository.Store(nUser)

		assert.NotNil(t, err)
		db.AssertCalled(t, "Read")
		db.AssertNotCalled(t, "Write", expect)
		db.AssertNotCalled(t, "Write")
	})

	t.Run("creating new with error in the read", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var nUser User = *NewUser("luiz", "lima", "luiz@gmail", 20, 1.7, true)
		expect := listExpect(t)

		db.On("Read").Return(expect, nil).Once()
		expect = append(expect, nUser)
		db.On("Write", expect).Return(errors.New("")).Once()
		_, err := repository.Store(nUser)

		assert.NotNil(t, err)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", expect)
	})

}

func TestRespositoryFindById(t *testing.T) {
	t.Run("Testing if FindById with Nothing inside", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var expected []User

		db.On("Read").Return(expected, nil).Once()
		user, err := repository.FindById("asdaq13")

		assert.NotNil(t, err)
		assert.Equal(t, User{}, user)
		db.AssertCalled(t, "Read")
	})
	t.Run("Testing if FindById with working fine", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expected := listExpect(t)

		db.On("Read").Return(expected, nil).Once()
		user, err := repository.FindById(expected[0].Id)

		assert.Nil(t, err)
		assert.Equal(t, expected[0], user)
		db.AssertCalled(t, "Read")
	})

	t.Run("Testing if FindById with working fine", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expected := listExpect(t)

		db.On("Read").Return(expected, nil).Once()
		user, err := repository.FindById("asdasdad")

		assert.NotNil(t, err)
		assert.Equal(t, User{}, user)
		db.AssertCalled(t, "Read")
	})

	t.Run("Testing if FindById with Read Error", func(t *testing.T) {
		db, repository := CreateRepository(t)

		db.On("Read").Return(nil, errors.New("adas")).Once()
		user, err := repository.FindById("asdasdad")

		assert.NotNil(t, err)
		assert.Equal(t, User{}, user)
		db.AssertCalled(t, "Read")
	})
}

func TestRespositoryFindByEmail(t *testing.T) {
	t.Run("Testing if FindByEmail with Nothing inside", func(t *testing.T) {
		db, repository := CreateRepository(t)
		var expected []User

		db.On("Read").Return(expected, nil).Once()
		user, err := repository.FindByEmail("asdaq13")

		assert.NotNil(t, err)
		assert.Equal(t, User{}, user)
		db.AssertCalled(t, "Read")
	})
	t.Run("Testing if FindByEmail with working fine", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expected := listExpect(t)

		db.On("Read").Return(expected, nil).Once()
		user, err := repository.FindByEmail(expected[0].Email)

		assert.Nil(t, err)
		assert.Equal(t, expected[0], user)
		db.AssertCalled(t, "Read")
	})

	t.Run("Testing if FindByEmail with working fine", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expected := listExpect(t)

		db.On("Read").Return(expected, nil).Once()
		user, err := repository.FindByEmail("asdasdad")

		assert.NotNil(t, err)
		assert.Equal(t, User{}, user)
		db.AssertCalled(t, "Read")
	})
	t.Run("Testing if findByEmail with Read Error", func(t *testing.T) {
		db, repository := CreateRepository(t)

		db.On("Read").Return(nil, errors.New("adas")).Once()
		user, err := repository.FindById("asdasdad")

		assert.NotNil(t, err)
		assert.Equal(t, User{}, user)
		db.AssertCalled(t, "Read")
	})
}

func TestRepositoryUpdate(t *testing.T) {
	t.Run("Testing Update nothing changes", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)
		expectUser := expect[0]
		deleted := slices.Clone(expect[1:])

		db.On("Read").Return(expect, nil).Once()
		db.On("Write", deleted).Return(nil).Once()
		db.On("Write", expect).Return(nil).Once()
		usr, err := repository.Update(expect[0], expect[0])

		assert.Nil(t, err)
		assert.Equal(t, expectUser, usr)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", deleted)
		db.AssertCalled(t, "Write", expect)
	})

	t.Run("Testing Update changes", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)
		oriUser := expect[0]
		expectUser := User{Id: "", Name: "pokemon", SourName: "milena", Email: "test@gmail", Age: 10}
		deleted := slices.Clone(expect[1:])

		db.On("Read").Return(expect, nil).Once()
		db.On("Write", deleted).Return(nil).Once()
		db.On("Write", expect).Return(nil).Once()
		usr, err := repository.Update(expect[0], expectUser)

		assert.Nil(t, err)
		expectUser.Id = oriUser.Id
		expectUser.Heigth = oriUser.Heigth
		expectUser.IsActive = oriUser.IsActive
		expectUser.CreateadAt = oriUser.CreateadAt
		assert.Equal(t, expectUser, usr)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", deleted)
		db.AssertCalled(t, "Write", expect)
	})

	t.Run("Testing Update changes Error on Fist Call", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)
		expectUser := User{Id: "", Name: "pokemon", SourName: "milena", Email: "test@gmail", Age: 10}
		deleted := slices.Clone(expect[1:])

		db.On("Read").Return(nil, errors.New("")).Once()
		// db.On("Write", deleted).Return(nil).Once()
		// db.On("Write", expect).Return(nil).Once()
		usr, err := repository.Update(expect[0], expectUser)

		assert.NotNil(t, err)
		assert.Equal(t, User{}, usr)
		db.AssertCalled(t, "Read")
		db.AssertNotCalled(t, "Write", deleted)
		db.AssertNotCalled(t, "Write", expect)
	})

	t.Run("Testing Update changes Error on Second Call", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)
		expectUser := User{Id: "", Name: "pokemon", SourName: "milena", Email: "test@gmail", Age: 10}
		deleted := slices.Clone(expect[1:])

		db.On("Read").Return(expect, nil).Once()
		db.On("Write", deleted).Return(errors.New("er")).Once()
		// db.On("Write", expect).Return(nil).Once()
		usr, err := repository.Update(expect[0], expectUser)

		assert.NotNil(t, err)
		assert.Equal(t, User{}, usr)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", deleted)
		db.AssertNotCalled(t, "Write", expect)
	})

	t.Run("Testing Update changes Error on Third Call", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)

		db.On("Read").Return(expect, nil).Once()
		db.On("Write", mock.Anything).Return(nil).Once()
		db.On("Write", mock.Anything).Return(errors.New("erasdjajd")).Once()
		usr, err := repository.Update(expect[0], expect[0])

		assert.NotNil(t, err)
		assert.Equal(t, User{}, usr)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", mock.Anything)
		db.AssertCalled(t, "Write", mock.Anything)
	})
}

func TestRepositoryDelete(t *testing.T) {
	t.Run("testing if deleted existing user", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)
		deleted := slices.Clone(expect[1:])

		db.On("Read").Return(expect, nil)
		db.On("Write", deleted).Return(nil)
		err := repository.Delete(expect[0])

		assert.Nil(t, err)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", deleted)
	})

	t.Run("testing if deleted not existing user", func(t *testing.T) {
		db, repository := CreateRepository(t)
		expect := listExpect(t)

		db.On("Read").Return(expect, nil)
		db.On("Write", expect).Return(nil)
		err := repository.Delete(*NewUser("luiz", "test", "lima@gmail.com", 20, 1.7, true))

		assert.Nil(t, err)
		db.AssertCalled(t, "Read")
		db.AssertCalled(t, "Write", expect)
	})

}

func TestRepositoryGetQueryParametersValid(t *testing.T) {
	_, repository := CreateRepository(t)
	res := repository.GetQueryParametersValids()
	assert.Equal(t, map[string]int{"name": 2, "sourname": 3, "email": 4, "age": 5, "height": 6, "isactive": 7}, res)
}

func TestReposioryNew(t *testing.T) {
	rep := NewRepository("test.json")

	assert.NotNil(t, rep)
	assert.IsType(t, &UserRepository{}, rep)
}
