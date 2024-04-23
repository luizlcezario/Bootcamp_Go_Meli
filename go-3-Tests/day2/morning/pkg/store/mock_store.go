package store

type MockStore struct {
	WasCalledRead  int
	WasCalledWrite int
	Data           interface{}
	Rt             error
}

func (m *MockStore) Read(data interface{}) (interface{}, error) {
	m.WasCalledRead++
	return m.Data, m.Rt
}

func (m *MockStore) Write(data interface{}) (interface{}, error) {
	m.WasCalledWrite++
	m.Data = data
	return data, m.Rt
}
