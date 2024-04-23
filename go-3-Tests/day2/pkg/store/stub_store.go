package store

type StubStore struct {
}

func (s *StubStore) Read(data interface{}) error {
	return nil
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}
