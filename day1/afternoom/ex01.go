package main

import (
	"fmt"
	"strings"
)

// nessa funcao eu separa todas as letras de uma string em uma lista e depois eu juntos elas colocando um espaco poderia ter colocado um \n ou qualquer outra string.

func preperWordToSpell(word string) (int, string) {
	result := strings.Split(word, "")
	return len(word), strings.Join(result, " ")
}

func main() {
	word := "basecamp"
	s, n := preperWordToSpell(word)
	fmt.Printf("Size of word: %d\nWord to Spell: %s\n", s, n)
}
