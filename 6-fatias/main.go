package main

import "fmt"

func main() {
	//slice, por baixo do pano é um array
	//ele tem um ponteiro apontando pro array, tamanho e capacidade
	//forma de declaração s := []int{}
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	//apaguei tudo o que tinha no meu slice desde a posição 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	//apaguei tudo o que tinha no meu slice após a posição 4
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	//ignorei o início, porém, também diminui a capacidade
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	//como aumento a capacidade do slice?
	//não preciso fazer nada, ele vai inserindo mais arrays (dobra o tamanho do slice)
	//tente iniciar o slice com o tamanho que tu imagina
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}
