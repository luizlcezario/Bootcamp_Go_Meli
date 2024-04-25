package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockStore struct {
	mock.Mock
}

func (m *MockStore) Read() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}

func (m *MockStore) Write(data interface{}) error {
	args := m.Called(data)
	return args.Error(0)
}
