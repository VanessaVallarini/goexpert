package main

func main() {
	//laço que normalmente usamos no go
	for i := 0; i < 10; i++ {
		println(i)
	}

	//também temos o range
	println("range: key and value")
	numeros := []string{"um", "dois", "três"}
	for k, v := range numeros {
		println(k, v)
	}

	println("range: only key")
	for k, _ := range numeros {
		println(k)
	}

	println("range: only value")
	for _, v := range numeros {
		println(v)
	}

	//for na pegada do while
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//loop infinito (consumir msg de uma fila)
	//for {
	//println("Hello, World!")
	//}

}
