package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Produtos.csv")
	if err != nil {
		panic("error ao tentar abrir o arquivo")
	}
	writer := csv.NewWriter(file)
	writer.Comma = ';'
	record := make([][]string, len(produtos)+1)
	for i, obj := range produtos {
		record[i+1] = []string{
			fmt.Sprint(obj.Id),
			fmt.Sprint(obj.Amount),
			fmt.Sprint(obj.Price)}
	}
	err = writer.WriteAll(record)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
}
