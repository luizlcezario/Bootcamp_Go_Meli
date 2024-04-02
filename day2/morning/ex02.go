package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func Sum(slice []int) (result int) {
	for _, a := range slice {
		result += a
	}
	return
}

func CalculateMedium(grades []int) int {
	if slices.Min(grades) < -1 {
		return 0
	}
	return Sum(grades) / len(grades)
}

func main() {

	for range 10 {
		len := rand.Intn(30)
		if len <= 0 {
			continue
		}
		slice := []int{}
		for range len {
			slice = append(slice, rand.Intn(10))
		}
		fmt.Printf("for the list of grades: %d the medium is: %d\n", slice, CalculateMedium(slice))
	}
}
