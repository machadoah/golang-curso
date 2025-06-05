package main

import "fmt"

func main() {
	numero := 10

	// if / else
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// if / init
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("POSITIVO")
	} else {
		fmt.Println("NEGATIVO")
	}

	// fmt.Println(outroNumero) -> nao existe fora do escopo condicional

}
