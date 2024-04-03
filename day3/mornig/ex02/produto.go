package main

import "reflect"

type Model interface {
	fieldsList() ([]string, int)
}

type Produto struct {
	Id     uint
	Price  float64
	Amount int
}

func (p Produto) fieldsList() ([]string, int) {
	val := reflect.ValueOf(p)
	fields := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fields[i] = val.Type().Field(i).Name
	}
	return fields, val.NumField()
}

var produtos []Model = []Model{
	&Produto{
		Id:     1,
		Price:  1000,
		Amount: 10,
	},
	&Produto{
		Id:     2,
		Price:  1100,
		Amount: 7,
	},
	&Produto{
		Id:     3,
		Price:  2000,
		Amount: 20,
	},
	&Produto{
		Id:     4,
		Price:  33300,
		Amount: 1,
	},
	&Produto{
		Id:     5,
		Price:  100,
		Amount: 200,
	},
}
