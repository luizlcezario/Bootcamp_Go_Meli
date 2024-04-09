package main

import (
	"errors"
	"fmt"
	"os"
)

type Reader struct {
	file string
}

func (r Reader) ReadFile() (string, error) {
	bytes, err := os.ReadFile(r.file)
	if err != nil {
		return "", errors.New("o arquivo indicado não foi encontrado ou está danificado ")
	}
	return string(bytes), nil
}

func main() {
	os := Reader{file: "customers.txt"}
	if vc, err := os.ReadFile(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(vc)
	}
}
