package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func (c Cliente) andou() {
	c.nome = "Vanessa Ichikawa"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func (c *Cliente) andou2() {
	c.nome = "Vanessa Ichikawa"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func (c Conta) simulador(valor int) int {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) simulador2(valor int) int {
	c.saldo += valor
	return c.saldo
}

//isso é muito usado em go, pois, eu crio uma conta e em qlq lugar que for usála estarei usando o seu endereço de memória
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	van := Cliente{
		nome: "Van",
	}

	van.andou()
	fmt.Printf("O valor da struct com nome %v\n", van.nome) //o nome só muda dentro da func andou()

	van.andou2()
	fmt.Printf("O valor da struct com nome %v\n", van.nome) //o nome muda dentro da func andou() e na struct van pq estou usando o endereço dee memória na func andou2()

	conta := Conta{
		saldo: 100,
	}
	conta.simulador(200)
	println(conta.saldo) //o resultado é 100 pq a alteração só vale dentro do contexto da func simulador()

	conta.simulador2(200)
	println(conta.saldo) //o resultado é 300 pq a alteração só feita no endereço de memória onde minha conta está armazenada

	conta2 := NewConta()
	conta2.simulador2(500)
	println(conta2.saldo)
}
