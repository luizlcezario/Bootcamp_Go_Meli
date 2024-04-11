package users

import (
	"encoding/json"
	"errors"
	"os"
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
		copy(usersCache, usersR)
		return usersR, nil
	}
}

func (m *MemoryUserRepository) Store(us User) (User, error) {
	users, err := m.GetAll()
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	users = append(users, us)
	jsonW, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	err = os.WriteFile(m.fileName, jsonW, 0644)
	if err != nil {
		return User{}, errors.New("failed to save in the file reversing the additing")
	}
	copy(usersCache, users)
	return us, nil
}

var usersCache []User = []User{}
