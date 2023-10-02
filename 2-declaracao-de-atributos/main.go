package main

//esse valor a não pode ser mudado pq é constante
const a = "Hello, World!"

//sintaxe: var nome tipo
//consigo alterar o seu valor
//não consigo mudar os tipos das variáveis, pois, é uma linguagem fortemente tipada
var b bool
var (
	c int
	d string
	//declarar a variável e já atribuir um valor
	f float64 = 1.2
)

func main() {
	//valores default
	println(a)
	println(b)
	println(c)
	println(d)
	println(f)

	b = true
	c = 1
	d = "Van"

	var e string
	println(a)
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	x()

	//criar, declarar e atribuir o valor de uma variável - usado em escopo local
	h := "h"
	//: só usa na criação
	h = "hh"
	println(h)
}

func x() {
	//não consegue acessar a variável e pq ela tem escopo local na func main
	//se eu declarar uma variável em escopo local eu devo usá-la, caso contrário dá erro
	var g bool
	println(g)
}
