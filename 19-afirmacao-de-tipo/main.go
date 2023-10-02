package main

import "fmt"

//typeAssertion (ok)

func main() {
	var minhaVar interface{} = "Van Ichikawa"
	println(minhaVar.(string)) //afirmando que essa var é uma string

	//se eu não usar o ok e der erro meu sistema da panic
	res, ok := minhaVar.(int)
	if !ok {
		fmt.Println("Erro na conversão!")
	} else {
		fmt.Println(res)
	}
}
