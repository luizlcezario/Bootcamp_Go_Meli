package main

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func NovoProduto(name string, price float64, amount int) *Produto {
	return &Produto{
		Nome:       name,
		Preco:      price,
		Quantidade: amount,
	}
}
