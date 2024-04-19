package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

type Reader struct {
	file string
}

type ErrorOpenFile struct{}

func (ErrorOpenFile) Error() string {
	return "o arquivo indicado não foi encontrado ou está danificado"
}

type Client struct {
	arquivo   string
	nome      string
	sobrenome string
	rg        string
	telefone  int
	endereco  string
}

func NewClient(nome, sobrenome, rg, endereco string, telefone int) *Client {
	return &Client{
		generatedId(nome + sobrenome + rg),
		nome,
		sobrenome,
		rg,
		telefone,
		endereco,
	}
}

func generatedId(s string) string {
	h := sha256.New()
	if _, err := h.Write([]byte(s)); err != nil {
		panic(err.Error())
	} else {
		bt := h.Sum(nil)
		return fmt.Sprintf("%x\n", bt)
	}
}

func (r Reader) SearchClient(rg string) *Client {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main the ReadFile", r)
		}
	}()
	if vc, err := r.ReadFile(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(vc)
		return NewClient(vc, vc, vc, vc, 10)
	}
}

func (r Reader) ReadFile() (string, error) {

	bytes, err := os.ReadFile(r.file)
	if err != nil {
		return "", ErrorOpenFile{}
	}
	return string(bytes), nil
}

func main() {
	os := Reader{file: "customers.txt"}
	client := os.SearchClient("12313")
	client2 := NewClient("luiz", "test", "23141", "rua", 123423123)
	fmt.Println(client, client2)
}
