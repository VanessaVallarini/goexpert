package main

//execute this file: go run main.go
import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	//carro.nome = "nome" -> carro.nome undefined (type matematica.Carro has no field or method nome)
	//fmt.Println(carro.nome) -> carro.nome undefined (type matematica.Carro has no field or method nome)
	fmt.Println(carro.Andar())
	fmt.Println("Resultado: ", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
