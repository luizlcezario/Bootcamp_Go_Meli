package main

type Ecommerce interface {
	Total() float64
	Adicionar(p Produto)
}

type Vender interface {
	X()
}
type Loja struct {
	products []Produto
}

func (l Loja) Total() float64 {
	result := 0.0
	for _, pr := range l.products {
		result += pr.CalcularCusto()
	}
	return result
}

func (l *Loja) Adicionar(p Produto) {
	l.products = append(l.products, p)
}

func NovaLoja() Ecommerce {
	return &Loja{}
}
