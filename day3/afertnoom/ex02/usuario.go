package main

type Usuario struct {
	Nome      string
	Sobrenome string
	Idade     int
	Email     string
	Password  string
	Produtos  []Produto
}

func NovoUsuario() *Usuario {
	return &Usuario{}
}

func (u *Usuario) SetName(name, sobrenome string) {
	u.Nome = name
	u.Sobrenome = sobrenome
}

func (u *Usuario) SetIdade(idade int) {
	u.Idade = idade
}

func (u *Usuario) SetEmail(email string) {
	u.Email = email
}

func (u *Usuario) AdicionarPoduto(produto *Produto) {
	u.Produtos = append(u.Produtos, *produto)
}

func (u *Usuario) RemoverProdutos() {
	u.Produtos = []Produto{}
}
