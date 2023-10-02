package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

//criando funções para structs
func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s\n", c.Nome, c.Idade, c.Ativo, c.Cidade)
}

func main() {
	van := Cliente{
		Nome:  "Van",
		Idade: 32,
		Ativo: true,
	}
	van.Cidade = "São Paulo"
	van.Desativar()
}
