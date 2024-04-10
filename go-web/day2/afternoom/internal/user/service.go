package users

import (
	"errors"
	"net/url"

	"github.com/luizlcezario/MelicampGoWeb/cmd/server/utils"
)

type Service interface {
	GetAllAndFilter(query url.Values) ([]User, error)
	FindById(id string) (User, error)
	Store(user User) (User, error)
}

type UserService struct {
	repository Repository
}

func NewService(r Repository) Service {
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

func (us *UserService) Store(user User) (User, error) {
	return us.repository.Store(user)
}
