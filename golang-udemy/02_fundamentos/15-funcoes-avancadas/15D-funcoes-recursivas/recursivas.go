package main

import "fmt"

func fib(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fib(posicao-2) + fib(posicao-1)
}

func main() {
	fmt.Println("Funçoes recursivas!")

	var posicao uint = 7

	fmt.Println(
		fib(posicao))
}
