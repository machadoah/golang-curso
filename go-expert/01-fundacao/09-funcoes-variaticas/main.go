package main

import "fmt"

func main() {
	fmt.Println(sum(10, 25, 50, 15))
}

// varios parametros
func sum(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
