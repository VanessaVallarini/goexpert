package main

import "fmt"

func main() {
	//array é uma variável de tamanho fixo e posso percorrer ele
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Printf("Valores array: \n [0] = %v \n [1] = %v \n [2] = %v \n", meuArray[0], meuArray[1], meuArray[len(meuArray)-1])
	fmt.Printf("Tamanho do array: %v \n", len(meuArray))

	//percorrendo o array
	for i, v := range meuArray {
		fmt.Printf("O valor do índice %d é: %d \n", i, v)
	}

}
