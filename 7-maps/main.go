package main

import (
	"fmt"
)

func main() {
	//maps é uma chave de key value e não tem ordenação
	salarios := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	fmt.Println(salarios["Wesley"])
	salarios["Van"] = 5000
	fmt.Println(salarios["Wesley"])

	//caso eu não queira inicializar valores posso usar as duas formas abaixo
	sal1 := make(map[string]int)
	sal2 := map[string]int{}

	sal1["Van1"] = 5000
	sal2["Van2"] = 5000

	fmt.Println(sal1)
	fmt.Println(sal2)

	//percorrendo o map
	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	//ignorando o nome ou valor - uso o underline (_)
	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
