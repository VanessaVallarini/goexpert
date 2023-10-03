package matematica

//esse pacote deve existir na pasta go/src da GOROOT
func Soma[T int | float64](a, b T) T {
	return a + b
}
