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

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome  string
	Ativo bool
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s\n", c.Nome, c.Idade, c.Ativo, c.Cidade)
}

func (e Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("Nome: %s, Ativo: %t\n", e.Nome, e.Ativo)
}

func main() {
	van := Cliente{
		Nome:  "Van",
		Idade: 32,
		Ativo: true,
	}
	van.Cidade = "São Paulo"
	Desativacao(van)

	minaEmpresa := Empresa{
		Nome:  "Empresa",
		Ativo: true,
	}
	Desativacao(minaEmpresa)
}
