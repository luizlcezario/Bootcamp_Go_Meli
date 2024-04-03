package main

import (
	"fmt"
)

type SalaryError struct{}

func (m SalaryError) Error() string {
	return "erro: O salário digitado não está dentro do valor mínimo"
}

func main() {
	salario := 1000

	if salario < 15000 {
		fmt.Println(SalaryError{}.Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
