package main

import "fmt"

type Employee struct {
	category string
	hours    int
}

func (e *Employee) CalculateSalary() (salary float64) {
	bonus := 1.0
	switch e.category {
	case "C":
		salary = float64(e.hours) * 1000
	case "B":
		salary = float64(e.hours) * 1500
		bonus = 1.2
	case "A":
		salary = float64(e.hours) * 1500
		bonus = 1.5
	}
	if e.hours > 160 {
		return salary * bonus
	}
	return
}

func main() {
	employe := []Employee{{category: "C", hours: 162}, {category: "B", hours: 176}, {category: "A", hours: 172}}
	for _, em := range employe {
		fmt.Println("Para o funcionario ", em)
		fmt.Println("O salario desse mes sera: ", em.CalculateSalary())
	}
}
