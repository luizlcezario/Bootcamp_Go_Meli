package main

import (
	"errors"
	"fmt"
)

func CalcSalary(hourValue float64, hours int) (float64, error) {
	result := hourValue * float64(hours)
	if hours < 80 {
		return 0, errors.Unwrap(fmt.Errorf("e2: %w", errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")))
	}
	if result > 20000 {
		return result * 1.2, nil
	}
	return result, nil
}

func main() {

	if result, err := CalcSalary(100, 10); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}
