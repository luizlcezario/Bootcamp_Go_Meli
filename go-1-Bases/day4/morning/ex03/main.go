package main

import (
	"fmt"
)

func main() {
	salario := 1000

	if salario < 15000 {
		fmt.Println(fmt.Errorf("erro: o mínimo tributável é 15.000 e o sálario informado é:%d", salario).Error())
	} else {
		fmt.Println("Necessário pagamento de imposto")
	}
}
