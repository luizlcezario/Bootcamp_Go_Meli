package store

type StubStore struct {
}

func (s *StubStore) Read(data interface{}) (interface{}, error) {
	return data, nil
}

func (s *StubStore) Write(data interface{}) (interface{}, error) {
	return data, nil
}
