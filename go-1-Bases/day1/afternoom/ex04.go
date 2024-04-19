package main

import "fmt"

func main() {
	// part 1: ajude a imprimir a idade de Benjamin.
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	fmt.Println(employees)
	fmt.Println("Benjamin tem:", employees["Benjamin"], "anos")
	// part 2: Saiba quantos de seus funcionários têm mais de 21 anos.
	ageGreaterthan := 0
	for _, value := range employees {
		if value > 21 {
			ageGreaterthan += 1
		}
	}
	fmt.Printf("tem %d funcionarios com mais de 21 anos\n", ageGreaterthan)
	// part 3: Adicione um novo funcionário à lista, chamado Federico, que tem 25 anos.
	employees["Federico"] = 25
	fmt.Println(employees)
	// part 4: Excluir Pedro do mapa.
	delete(employees, "Pedro")
	fmt.Println(employees)
}
