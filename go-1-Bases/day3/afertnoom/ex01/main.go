package main

import "fmt"

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Password  string
}

func NovoUsuario() *Usuario {
	return &Usuario{}
}

func (u *Usuario) setName(name, sobrenome string) {
	u.Nome = name
	u.Sobrenome = sobrenome
}

func (u *Usuario) setIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) setEmail(email string) {
	u.Email = email
}

func main() {
	usuario := NovoUsuario()
	usuario.setIdade(20)
	usuario.setName("luiz", "cezario")
	usuario.setEmail("luizlcezario@gmail.com")
	fmt.Println(usuario)
}
