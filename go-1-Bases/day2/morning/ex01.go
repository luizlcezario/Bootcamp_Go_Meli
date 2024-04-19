package main

import (
	"fmt"
	"math/rand"
)

func CalculateTaxSalary(salary float64) float64 {
	switch {
	case salary > 150000:
		return salary * 0.27
	case salary > 50000:
		return salary * 0.17
	default:
		return 0
	}
}

func main() {
	for range 100 {
		salary := rand.Float64() * 100000
		fmt.Printf("Este e o valor do desconto: %0.2f Para um usuario que recebe: %0.2f\n", CalculateTaxSalary(salary), salary)
	}
}
