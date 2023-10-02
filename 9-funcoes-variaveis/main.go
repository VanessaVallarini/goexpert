package main

import "fmt"

func main() {
	fmt.Println(sum(1, 50, 90, 40, 70, 20, 30, 45, 400))
}

//quero somar uma quantidade indefinida de números
func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
