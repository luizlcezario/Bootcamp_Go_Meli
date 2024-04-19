package main

type Sizes int

const (
	Undefined Sizes = iota
	Pequeno
	Medio
	Grande
)

type Produto interface {
	CalcularCusto() float64
}

type produto struct {
	name          string
	price         float64
	sizeOfProduct Sizes
}

// CalcularCusto implements Produto.

func NovoProduto(name string, price float64, tipe Sizes) produto {
	return produto{
		name:          name,
		price:         price,
		sizeOfProduct: tipe,
	}
}

func (p produto) CalcularCusto() float64 {
	switch p.sizeOfProduct {
	case Medio:
		return p.price * 1.03
	case Grande:
		return (p.price * 1.06) + 2500
	default:
		return p.price
	}
}
