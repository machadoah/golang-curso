package main

import "fmt"

func soma(numeros ...int) int {
	// numeros é representado por um slice
	total := 0
	for _, num := range numeros {
		total += num
	}

	return total
}

func escrever(texto string, numeros ...int) {
	total := 0
	for _, num := range numeros {
		total += num
	}

	println(texto, total)
}

func main() {
	fmt.Println(soma(1, 3, 345, 5, 2, 545, 77, 54, 32))

	escrever("A", 1, 3, 5, 8)
}
