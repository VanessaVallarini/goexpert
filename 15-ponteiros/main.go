package main

func main() {
	//Memória -> Endereço -> Valor
	a := 10
	println(&a) //traz o endereço de memória de a

	var ponteiro *int = &a //pegando o endereço de a
	println(ponteiro)

	*ponteiro = 20
	println(a) //vale 20 pq ponteiro aponta para a

	b := &a
	println(*b) //qual valor está guardado nessa memória de b

	*b = 30
	//a e b valem 30
	println(a)
	println(*b)
}
