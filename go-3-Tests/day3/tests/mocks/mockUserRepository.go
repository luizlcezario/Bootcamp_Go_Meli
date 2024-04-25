package mocks

import (
	. "github.com/luizlcezario/MelicampGoWeb/internal/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}
func (m *MockUserRepository) FindById(id string) (User, error) {
	args := m.Called()
	return args.Get(0).(User), args.Error(1)
}
func (m *MockUserRepository) FindByEmail(email string) (User, error) {
	args := m.Called()
	return args.Get(0).(User), args.Error(1)
}
func (m *MockUserRepository) Store(us User) (User, error) {
	args := m.Called()
	return args.Get(0).(User), args.Error(1)
}
func (m *MockUserRepository) Update(dst User, src User) (User, error) {
	args := m.Called()
	return args.Get(0).(User), args.Error(1)
}
func (m *MockUserRepository) Delete(us User) error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockUserRepository) GetQueryParametersValids() map[string]int {
	args := m.Called()
	return args.Get(0).(map[string]int)
}
