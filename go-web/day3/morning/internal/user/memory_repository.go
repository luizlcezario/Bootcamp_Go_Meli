package users

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
)

type MemoryUserRepository struct {
	fileName string
}

func NewRepository(file string) Repository {
	return &MemoryUserRepository{
		fileName: file,
	}
}

func (*MemoryUserRepository) GetQueryParametersValids() map[string]int {
	return map[string]int{"name": 2, "sourname": 3, "email": 4, "age": 5, "height": 6, "isactive": 7}
}

func (re *MemoryUserRepository) FindById(id string) (User, error) {
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

func (re *MemoryUserRepository) FindByEmail(email string) (User, error) {
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

func (m *MemoryUserRepository) GetAll() ([]User, error) {
	if len(usersCache) != 0 {
		return usersCache, nil
	} else {
		jsonR, err := os.ReadFile(m.fileName)
		if err != nil {
			return []User{}, errors.New("not possible to open the file to read the data")
		}
		var usersR []User
		err = json.Unmarshal(jsonR, &usersR)
		if err != nil {
			return []User{}, errors.New("not possible to unmarshal the json of the file to the struct")
		}
		usersCache = make([]User, len(usersR))
		copy(usersCache, usersR)
		return usersR, nil
	}
}

func (m *MemoryUserRepository) Store(us User) (User, error) {
	user, err := m.GetAll()
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	fmt.Println(user)
	user = append(user, us)
	jsonW, err := json.MarshalIndent(user, "", " ")
	fmt.Println(user)
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	err = os.WriteFile(m.fileName, jsonW, 0644)
	if err != nil {
		return User{}, errors.New("failed to save in the file reversing the additing")
	}
	fmt.Println(user)
	return us, nil
}

func (e *MemoryUserRepository) Delete(us User) error {
	usersCache = slices.DeleteFunc(usersCache, func(e User) bool {
		return (e == us)
	})
	jsonW, err := json.MarshalIndent(usersCache, "", " ")
	if err != nil {
		return errors.New("failed to marshal the users list")
	}
	err = os.WriteFile(e.fileName, jsonW, 0644)
	if err != nil {
		return errors.New("failed to save in the file reversing the additing")
	}
	return nil
}

func (e *MemoryUserRepository) Update(dst, src User) User {
	fmt.Println(len(usersCache))
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
	fmt.Println(len(usersCache))
	return dst
}

var usersCache []User = []User{}
