package main

func main() {
	a := 1
	b := 2

	//OPERADORES LÓGICOS
	//>= maior igual
	//<= menor igual
	//&& e
	//|| ou
	//== igual

	//IF ELSE
	if a > b {
		println(a)
	} else {
		println(b)
	} //não temos elseif no go

	//switch
	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("c")
	}

}
