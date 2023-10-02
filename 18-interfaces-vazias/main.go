package main

import "fmt"

//devemos usar isso com muita moderação
//generics pode substituir a necessidade de usar interfaces vazias

func main() {
	//interface aceita qualquer coisa
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
