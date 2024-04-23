package store

import (
	"encoding/json"
	"errors"
	"os"
)

type FileStore struct {
	fileName string
}

func (f *FileStore) Read(data interface{}) (interface{}, error) {
	jsonR, err := os.ReadFile(f.fileName)
	if err != nil {
		return nil, errors.New("not possible to open the file to read the data")
	}
	return data, json.Unmarshal(jsonR, data)
}

func (f *FileStore) Write(data interface{}) (interface{}, error) {
	jsonW, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, errors.New("failed to marshal the data list")
	}
	return data, os.WriteFile(f.fileName, jsonW, 0644)
}
