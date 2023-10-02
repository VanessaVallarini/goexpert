package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum1(1, 2))
	fmt.Println(sum2(1, 2))
	fmt.Println(sum3(1, 2))

	valor, err := sum4(1, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	valor, err = sum4(1, 50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)
}

func sum1(a int, b int) int {
	return a + b
}

func sum2(a, b int) int {
	return a + b
}

func sum3(a, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

//esse tipo de função é muito usada para retornar o valor e o erro, pq go não tem os try cath
func sum4(a, b int) (int, error) {
	if a+b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
