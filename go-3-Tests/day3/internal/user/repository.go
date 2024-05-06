package users

import (
	"errors"
	"slices"

	. "github.com/luizlcezario/MelicampGoWeb/internal/domain"
	"github.com/luizlcezario/MelicampGoWeb/pkg/store"
)

type Repository interface {
	GetAll() ([]User, error)
	FindById(id string) (User, error)
	FindByEmail(email string) (User, error)
	Store(us User) (User, error)
	Update(dst User, src User) (User, error)
	Delete(us User) error
	GetQueryParametersValids() map[string]int
}

type UserRepository struct {
	db         store.Store
	usersCache []User
}

func NewRepository(db store.Store) Repository {
	return &UserRepository{
		db: store.NewStore[[]User](store.FileType, "user.json"),
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
	if len(m.usersCache) != 0 {
		return m.usersCache, nil
	}
	if user, err := m.db.Read(); err != nil {
		return m.usersCache, err
	} else {
		m.usersCache = make([]User, len(user.([]User)))
		copy(m.usersCache, user.([]User))
		return user.([]User), nil
	}
}

func (m *UserRepository) Store(us User) (User, error) {
	user, err := m.GetAll()
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	m.usersCache = append(user, us)
	if err := m.db.Write(m.usersCache); err != nil {
		return User{}, err
	}
	return us, nil
}

func (e *UserRepository) Delete(us User) error {
	user, err := e.GetAll()
	if err != nil {
		return err
	}
	e.usersCache = slices.DeleteFunc(user, func(e User) bool {
		return (e == us)
	})
	if err := e.db.Write(e.usersCache); err != nil {
		return err
	}
	return nil
}

func (e *UserRepository) Update(dst, src User) (User, error) {
	if err := e.Delete(dst); err != nil {
		return User{}, err
	}
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
	if _, err := e.Store(dst); err != nil {
		return User{}, err
	}
	return dst, nil
}
