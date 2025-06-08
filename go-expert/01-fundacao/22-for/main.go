package main

import "fmt"

func main() {
	// Go tem somente FOR

	// itera de 0 a 9
	fmt.Println("Maneira antiga:")
	for i := 0; i < 10; i++ {
		println(i)
	}

	// itera de 0 a 9
	fmt.Println("Maneira moderna:")
	for i := range 10 {
		println(i)
	}

	println("Percorrendo um slice com chace")
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	// simulando um while
	println("simulando um while")
	i := 0

	for i < 10 {
		println(i)
		i++
	}

	println("loop infito")

	for {
	}
}
