package main

import "fmt"

type Endereco struct {
	Logadouro string
	Numero    int
	Cidade    string
	Estado    string
}

//Compondo Endereço em Cliente
type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

//Diferente de compor
type Cliente2 struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func main() {
	van := Cliente{
		Nome:  "Van",
		Idade: 32,
		Ativo: true,
	}
	van.Cidade = "São Paulo" //posso usar van.Endereco.Cidade, porém, dá na mesma :)
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s\n", van.Nome, van.Idade, van.Ativo, van.Cidade)

	van2 := Cliente2{
		Nome:  "Van",
		Idade: 32,
		Ativo: true,
	}
	van2.Address.Cidade = "São Paulo" //não posso usar van.Cidade :(
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", van2.Nome, van2.Idade, van2.Ativo, van2.Address.Cidade)
}
