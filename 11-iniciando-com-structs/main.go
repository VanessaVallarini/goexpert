package main

import "fmt"

//struct imaginamos que é um tipo de uma classe
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	//go não é orientada a objetos (herança via composição - uma struct dentro de outra)
	van := Cliente{
		Nome:  "Van",
		Idade: 32,
		Ativo: true,
	}

	van.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", van.Nome, van.Idade, van.Ativo)
}
