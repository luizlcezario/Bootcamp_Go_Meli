package store

import (
	"encoding/json"
	"errors"
	"os"
)

type FileStore[T interface{}] struct {
	fileName string
}

func (f *FileStore[T]) Read() (interface{}, error) {
	jsonR, err := os.ReadFile(f.fileName)
	if err != nil {
		return nil, errors.New("not possible to open the file to read the data")
	}
	var data T
	err = json.Unmarshal(jsonR, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *FileStore[T]) Write(data interface{}) error {
	jsonW, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return errors.New("failed to marshal the data list")
	}
	return os.WriteFile(f.fileName, jsonW, 0644)
}
