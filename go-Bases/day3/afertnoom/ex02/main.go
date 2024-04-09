package main

import "fmt"

func main() {
	usuario := NovoUsuario()
	usuario.SetIdade(20)
	usuario.SetName("luiz", "cezario")
	usuario.SetEmail("luizlcezario@gmail.com")
	usuario.AdicionarPoduto(NovoProduto("geladeira", 4500, 1))
	usuario.AdicionarPoduto(NovoProduto("xbox", 3000, 3))
	usuario.AdicionarPoduto(NovoProduto("microondas", 500, 5))
	fmt.Println(usuario)
	usuario.RemoverProdutos()
	fmt.Println(usuario)

}
