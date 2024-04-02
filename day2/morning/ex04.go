package main

import (
	"errors"
	"fmt"
	"slices"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func sum(number ...int) (result int) {
	for _, number := range number {
		result += number
	}
	return
}

func operation(str string) (func(number ...int) int, error) {
	switch str {
	case minimum:
		return func(number ...int) int {
			return slices.Min(number)
		}, nil
	case maximum:
		return func(number ...int) int {
			return slices.Max(number)
		}, nil
	case average:
		return func(number ...int) int {
			return sum(number...) / len(number)
		}, nil
	default:
		return nil, errors.New("Calculo nao esta definido")
	}
}

func main() {

	minhaFunc, err := operation(minimum)
	averageFunc, err := operation(average)
	maxFunc, err := operation(maximum)
	if err != nil {
		panic(err.Error())
	} else {
		minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)

		averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

		maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

		fmt.Println("minValue: ", minValue, " averageValue: ", averageValue, " maxvalue: ", maxValue)
	}
}
