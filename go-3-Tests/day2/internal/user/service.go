package users

import (
	"errors"
	"net/url"

	"github.com/luizlcezario/MelicampGoWeb/cmd/server/utils"
)

type Service interface {
	GetAllAndFilter(query url.Values) ([]User, error)
	FindById(id string) (User, error)
	Store(name, sourName, email string, age int, heigth float32) (User, error)
	UpdateUser(id, name, sourName, email string, age int, heigth float32) (User, error)
	Delete(id string) error
}

type UserService struct {
	repository Repository
}

func NewService(r Repository) Service {
	usersCache = []User{}
	return &UserService{
		repository: r,
	}
}

func (us *UserService) FindById(id string) (User, error) {
	return us.repository.FindById(id)

}

func (us *UserService) GetAllAndFilter(query url.Values) ([]User, error) {
	result, err := us.repository.GetAll()

	if err != nil {
		return []User{}, err
	}
	tags := us.repository.GetQueryParametersValids()
	for key, i := range query {
		if field := tags[key]; field == 0 {
			return []User{}, errors.New("please send just query parameters valids")
		} else {
			result = utils.FilterSliceByQueryListValue(field, i, result)
		}
	}
	return result, nil
}

func (us *UserService) Store(name, sourName, email string, age int, height float32) (User, error) {
	if _, err := us.repository.FindByEmail(email); err != nil {
		return us.repository.Store(*NewUser(name, sourName, email, age, height, true))
	}
	return User{}, errors.New("already Exists")
}

func (us *UserService) UpdateUser(id, name, sourName, email string, age int, heigth float32) (User, error) {
	if user, err := us.repository.FindById(id); err != nil {
		return User{}, errors.New("User not found")
	} else {
		newUser := User{
			Name:     name,
			SourName: sourName,
			Email:    email,
			Age:      age,
			Heigth:   heigth,
		}
		usr := us.repository.Update(user, newUser)
		return usr, nil
	}
}

func (us *UserService) Delete(id string) error {
	if user, err := us.repository.FindById(id); err != nil {
		return errors.New("User not found")
	} else {
		return us.repository.Delete(user)
	}
}
