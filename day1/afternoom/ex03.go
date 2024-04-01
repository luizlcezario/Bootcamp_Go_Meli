package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1. Dependendo do número, imprima o mês correspondente em texto.
func dateNumToString(date string) string {
	mesesById := map[int]string{
		1:  "janeiro",
		2:  "fevereiro",
		3:  "março",
		4:  "abril",
		5:  "maio",
		6:  "junho",
		7:  "julho",
		8:  "agosto",
		9:  "setembro",
		10: "outubro",
		11: "novembro",
		12: "dezembro",
	}
	numbersStr := strings.Split(date, "/")
	numbers := [3]int{0, 0, 0}
	for i, str := range numbersStr {
		if num, err := strconv.Atoi(str); err != nil {
			panic("Erro no strconvAtoi")
		} else {
			numbers[i] = num
		}
	}
	return fmt.Sprintf("Dia %d de %s de %d\n", numbers[0], mesesById[numbers[1]], numbers[2])
}

// 2. Ocorre a você que isso pode ser resolvido de maneiras diferentes? Qual você escolheria e porquê?
// Sim acredito que com a lista de messes em um map seja a melhor pois o acesso a Key e O(1) time complexy
func main() {
	date := "01/10/2000"
	fmt.Printf(dateNumToString(date))
}
