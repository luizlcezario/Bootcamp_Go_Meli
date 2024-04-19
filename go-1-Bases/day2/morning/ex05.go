package main

import (
	"errors"
	"fmt"
)

const (
	dog       = "dog"
	cat       = "cat"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func Animal(animal string) (func(int) float64, error) {
	value := 0.0
	switch animal {
	case dog:
		value = 10
	case cat:
		value = 5
	case hamster:
		value = 0.250
	case tarantula:
		value = 0.150
	default:
		return nil, errors.New("Error este animal ainda nao pode ser mantido no abrigo")
	}
	return func(i int) float64 {
		return float64(i) * value
	}, nil
}

func main() {
	animalDog, err := Animal(dog)
	animalCat, err := Animal(cat)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)
	if err != nil {
		panic(err.Error())
	}
	var amount float64
	amount += animalDog(5)
	amount += animalCat(8)
	amount += animalHamster(12)
	amount += animalTarantula(13)

	fmt.Println("no total precisamos de ", amount, " kilos de comida")
}
