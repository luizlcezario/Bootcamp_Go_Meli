package store

import (
	"bytes"
	"encoding/gob"
)

type MockStore struct {
	WasCalledRead  bool
	WasCalledWrite bool
	Data           interface{}
	Rt             error
}

func (m *MockStore) Read(data interface{}) error {
	m.WasCalledRead = true
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(m.Data)
	gob.NewDecoder(&buf).Decode(data)
	return m.Rt
}

func (m *MockStore) Write(data interface{}) error {
	m.WasCalledWrite = true
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(data)
	gob.NewDecoder(&buf).Decode(m.Data)
	return m.Rt
}
