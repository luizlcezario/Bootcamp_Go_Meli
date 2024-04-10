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
		usersCache = usersR
		return usersCache, nil
	}
}

func (m *MemoryUserRepository) Store(us User) (User, error) {
	users := append(usersCache, us)
	jsonW, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return User{}, errors.New("failed to marshal the users list")
	}
	err = os.WriteFile(m.fileName, jsonW, 0644)
	if err != nil {
		return User{}, errors.New("failed to save in the file reversing the additing")
	}
	usersCache = users
	return us, nil
}

var usersCache []User = []User{
	*NewUser("luiz", "lima", "luiz@gmail", 20, 1.76, true),
	*NewUser("milena", "carecho", "care@gmail", 22, 1.56, true),
	*NewUser("luan", "cardoso", "lcar@gmail", 23, 1.67, false),
	*NewUser("felipe", "lima", "fel@gmail", 10, 1.90, true),
}
