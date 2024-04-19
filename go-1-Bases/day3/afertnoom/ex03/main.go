package main

import (
	"fmt"
)

func main() {
	priceFinal := make(chan float64)

	produtos := []preço{
		NovoProduto("geladeira", 4500, 1),
		NovoProduto("xbox", 3000, 3),
		NovoProduto("microondas", 500, 5),
	}
	servicos := []preço{
		NovoServico("tester", 123.23, 24),
		NovoServico("tester", 982.3, 2.2),
		NovoServico("tester", 100, 16.8),
	}
	manutencao := []preço{
		NovoManutencao("tester", 123.23),
		NovoManutencao("tester", 982.3),
		NovoManutencao("tester", 100),
	}
	go PriceAllFinal(produtos, priceFinal)
	go PriceAllFinal(servicos, priceFinal)
	go PriceAllFinal(manutencao, priceFinal)
	final := <-priceFinal + <-priceFinal + <-priceFinal
	fmt.Printf("%.2f\n", final)
	close(priceFinal)
}
