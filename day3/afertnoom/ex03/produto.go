package main

type preço interface {
	Price() float64
}

func PriceAllFinal(objs []preço, result chan float64) {
	res := 0.0
	for _, obj := range objs {
		res += obj.Price()
	}
	result <- (res)
}

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

func (p *Produto) Price() float64 {
	return p.Preco * float64(p.Quantidade)
}
