package main

import "fmt"

func SomaInt(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

//generics
func SomaGenerics1[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type Number interface {
	int | float64
}

//generics
func SomaGenerics2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

type MyNumber int

type Number2 interface {
	~int | float64
}

//generics
func SomaGenerics3[T Number2](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

//comparando a igualdade de generics
//tem outro pacote que compara mais coisas? maior, menor... chama-se constraints
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	mInt := map[string]int{"Van": 1000, "João": 2000, "Maria": 3000}
	fmt.Printf("Soma int: %v\n", SomaInt(mInt))

	mFloat := map[string]float64{"Van": 100.20, "João": 200.30, "Maria": 300.50}
	fmt.Printf("Soma float: %v\n", SomaFloat(mFloat))

	fmt.Println()
	fmt.Printf("Soma Generics1 int (T): %v\n", SomaGenerics1(mInt))
	fmt.Printf("Soma Generics1 float (T): %v\n", SomaGenerics1(mFloat))

	fmt.Println()
	fmt.Printf("Soma Generics2 int (interface number): %v\n", SomaGenerics2(mInt))
	fmt.Printf("Soma Generics2 float (interface number): %v\n", SomaGenerics2(mFloat))

	fmt.Println()
	//usando um type para criar meu map e passando ele para generics
	mNumber2 := map[string]MyNumber{"Van": 1000, "João": 2000, "Maria": 3000}
	fmt.Printf("Soma Generics3 int (interface number que aceita tipo MyNumber): %v\n", SomaGenerics3(mNumber2))

	fmt.Println()
	fmt.Printf("Comparando a igualdade de generics: %v e %v são iguais? %v", 10, 10, Compara(10, 10))
}
