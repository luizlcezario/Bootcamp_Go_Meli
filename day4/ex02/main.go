package main

import (
	"errors"
	"fmt"
)

func main() {
	salario := 1000

	if salario < 15000 {
		fmt.Println(errors.New("erro: O salário digitado não está dentro do valor mínimo").Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
