package main

import "fmt"

func main() {
	loja := NovaLoja()
	loja.Adicionar(NovoProduto("xbox", 2500, Medio))
	fmt.Printf("Total da Loja: R$%0.2f\n", loja.Total())
	loja.Adicionar(NovoProduto("ps4", 2000.34, Medio))
	loja.Adicionar(NovoProduto("Tv 60pl", 4559.90, Grande))
	fmt.Printf("Total da Loja: R$%0.2f\n", loja.Total())
	loja.Adicionar(NovoProduto("secador", 120, Pequeno))
	fmt.Printf("Total da Loja: R$%0.2f\n", loja.Total())

}
