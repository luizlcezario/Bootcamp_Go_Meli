package main

import (
	"errors"
	"fmt"
)

// criei uma struct pra representar um client

type Client struct {
	name        string
	age         int
	isWorking   bool
	timeWorking int // in month
	salary      int
}

// se o client for elegivel para o emprestimo a funcao nao retorna nenhum error e o valor do juros na variavel do tipo float

func offerLoan(client Client) (float32, error) {
	if client.age < 22 {
		return 0.0, errors.New("este não esta elegivel para o emprestivo, devido a idade!")
	} else if !client.isWorking {
		return 0.0, errors.New("este não esta elegivel para o emprestivo, pois nao esta trabalhando!")
	} else if client.timeWorking < 12 {
		return 0.0, errors.New("este não esta elegivel para o emprestivo, poir trabalha a menos de um ano!")
	} else if client.salary >= 100000 {
		return 0.0, nil
	}
	return 1.4, nil
}

func main() {
	client := Client{
		name:        "luiz",
		age:         23,
		isWorking:   true,
		timeWorking: 15,
		salary:      1000000,
	}
	if juros, err := offerLoan(client); err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Este client é elegivel para um emprestimo, sua taxa de juros sera: %0.2f%% por mês!\n", juros)
	}

}
