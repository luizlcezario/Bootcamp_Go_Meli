package main

type produto struct {
	Id     uint
	Price  float64
	Amount int
}

var produtos []produto = []produto{
	{
		Id:     1,
		Price:  1000,
		Amount: 10,
	},
	{
		Id:     2,
		Price:  1100,
		Amount: 7,
	},
	{
		Id:     3,
		Price:  2000,
		Amount: 20,
	},
	{
		Id:     4,
		Price:  33300,
		Amount: 1,
	},
	{
		Id:     5,
		Price:  100,
		Amount: 200,
	},
}
