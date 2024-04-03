package main

type Manutencao struct {
	Nome   string
	PriceV float64
}

func NovoManutencao(name string, price float64) *Manutencao {
	return &Manutencao{
		Nome:   name,
		PriceV: price,
	}
}

func (m *Manutencao) Price() float64 {
	return m.PriceV
}
