package main

type Servico struct {
	Nome      string
	PriceV    float64
	workHours float64
}

func NovoServico(name string, price float64, workHours float64) *Servico {
	return &Servico{
		Nome:      name,
		PriceV:    price,
		workHours: workHours,
	}
}

func (s *Servico) Price() float64 {
	if s.workHours < 0.5 {
		return s.PriceV * 0.5
	}
	return s.PriceV * s.workHours
}
