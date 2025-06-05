package main

import "fmt"

func closure() func() {
	texto := "---------- DENTRO DA FUNCAO CLOSURE"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	texto := "---------- DENTRO DA FUNCAO MAIN"
	fmt.Println(texto)

	novaFuncao := closure()

	novaFuncao()
}
