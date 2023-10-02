package main

//USO QUANDO QUERO FAZER UM CÁLCULO E RETORNAR O RESULTADO
//faz uma cópia das minhas variáveis 1 e 2
//se eu alterar o valor de a, essa alteração só é refletida dentro dessa func
func soma(a, b int) int {
	a = 50
	println(a)
	return a + b
}

//USO QUANDO PRECISO TRABALHAR COM VALORES MUTÁVEIS
//usando o endereço de memória
func soma2(a, b *int) int {
	*a = 50
	println(*a)
	return *a + *b
}

func main() {
	minaVar1 := 10
	minaVar2 := 20
	println("CÓPIA")
	println(minaVar1)
	println(soma(minaVar1, minaVar2))

	println("MEMÓRIA")
	println(minaVar1)
	println(soma2(&minaVar1, &minaVar2))
}
