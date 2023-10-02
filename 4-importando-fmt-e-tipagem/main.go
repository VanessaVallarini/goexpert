package main

//como imprimir um valor para saber o tipo da variável
//formatando a impressão no fmt

import "fmt"

type ID int

func main() {
	a := 1.2
	var b ID = 1

	fmt.Printf("O tipo de A é %T\n", a)  //T traz o tipo da variável
	fmt.Printf("O tipo de B é %T\n", b)  //T traz o tipo da variável
	fmt.Printf("O valor de A é %v\n", a) //T traz o valor da variável
	fmt.Printf("O valor de B é %v\n", b) //T traz o valor da variável
}
