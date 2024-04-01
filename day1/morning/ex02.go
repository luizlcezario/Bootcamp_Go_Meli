package main

import "fmt"

func main() {
	var temperatura int = 26
	var umidade int = 57
	var pressao float64 = 1016.3

	fmt.Printf("temperatura:%d °\numidade: %d%% \npressao: %0.2fmb\n", temperatura, umidade, pressao)
}
