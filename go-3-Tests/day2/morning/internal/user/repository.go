package users

import (
	"errors"
	"slices"

	"github.com/luizlcezario/MelicampGoWeb/pkg/store"
)

type Repository interface {
	GetAll() ([]User, error)
	FindById(id string) (User, error)
	FindByEmail(email string) (User, error)
	Store(us User) (User, error)
	Update(dst User, src User) User
	Delete(us User) error
	GetQueryParametersValids() map[string]int
}

type UserRepository struct {
	db store.Store
}

func NewRepository(file string) Repository {
	return &UserRepository{
		db: store.NewStore(store.FileType, "user.json"),
	}
}

func (*UserRepository) GetQueryParametersValids() map[string]int {
	return map[string]int{"name": 2, "sourname": 3, "email": 4, "age": 5, "height": 6, "isactive": 7}
}

func (re *UserRepository) FindById(id string) (User, error) {
	if users, err := re.GetAll(); err != nil {
		return User{}, err
	} else {
		for _, e := range users {
			if e.Id == id {
				return e, nil
			}
		}
		return User{}, errors.New("Nao foi possivel achar o usuario com o Id: " + id)
	}
}

func (re *UserRepository) FindByEmail(email string) (User, error) {
	if users, err := re.GetAll(); err != nil {
		return User{}, err
	} else {
		for _, e := range users {
			if e.Email == email {
				return e, nil
			}
		}
		return User{}, errors.New("Nao foi possivel achar o usuario com o Email: " + email)
	}
}

func (m *UserRepository) GetAll() ([]User, error) {
	if len(usersCache) != 0 {
		return usersCache, nil
	}
	var user []User
	if _, err := m.db.Read(&user); err != nil {
		return usersCache, err
	} else {
		usersCache = make([]User, len(user))
		copy(usersCache, user)
		return user, nil
	}
}

func (m *UserRepository) Store(us User) (User, error) {
	user, err := m.GetAll()
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	usersCache = append(user, us)
	if _, err := m.db.Write(usersCache); err != nil {
		return User{}, err
	}
	return us, nil
}

func (e *UserRepository) Delete(us User) error {
	usersCache = slices.DeleteFunc(usersCache, func(e User) bool {
		return (e == us)
	})
	if _, err := e.db.Write(usersCache); err != nil {
		return err
	}
	return nil
}

func (e *UserRepository) Update(dst, src User) User {
	e.Delete(dst)
	if src.Name != "" {
		dst.Name = src.Name
	}
	if src.SourName != "" {
		dst.SourName = src.SourName
	}
	if src.Email != "" {
		dst.Email = src.Email
	}
	if src.Age != 0 {
		dst.Age = src.Age
	}
	if src.Heigth != 0.0 {
		dst.Heigth = src.Heigth
	}
	e.Store(dst)
	return dst
}

var usersCache []User = []User{}
