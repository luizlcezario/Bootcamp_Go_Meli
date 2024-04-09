package main

import "fmt"

type Aluno struct {
	nome      string
	sobrenome string
	rg        string
	data      string
}

func NewAluno() *Aluno {
	return &Aluno{}
}

func (a *Aluno) detalhamento() {
	fmt.Println("Nome: ", a.nome, " ", a.sobrenome)
	fmt.Println("RG: ", a.rg)
	fmt.Println("Data da entrada: ", a.data)
}

func main() {
	aluno := NewAluno()
	aluno.nome = "luiz"
	aluno.sobrenome = "cezario"
	aluno.rg = "34702838857"
	aluno.data = "26/10/2000"
	aluno2 := Aluno{
		nome:      "luan",
		sobrenome: "sobreposicao",
		rg:        "23214512321",
		data:      "10/10/2012",
	}
	aluno.detalhamento()
	aluno2.detalhamento()
}
