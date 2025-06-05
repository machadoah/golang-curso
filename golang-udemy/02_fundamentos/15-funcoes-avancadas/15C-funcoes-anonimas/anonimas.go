package main

import "fmt"

func main() {
	func() {
		fmt.Println("Função anônima")
	}()

	func(texto string) {
		fmt.Println(texto)
	}("Função anônima 2")

	retorno := func() string {
		return "Retorno"
	}()

	fmt.Println(retorno)
}
