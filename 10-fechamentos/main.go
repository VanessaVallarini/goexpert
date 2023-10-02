package main

import "fmt"

func main() {

	// uma função que é executada para retornar outro valor (* 2)
	total := func() int {
		return sum(1, 50, 90, 40, 70, 20, 30, 45, 400) * 2
	}()

	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
