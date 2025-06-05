package main

import "fmt"

func alunoIsAprovado(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	defer fmt.Println("Média foi calculada:", media)

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	if alunoIsAprovado(3, 3) == true {
		fmt.Println("Aprovado! 🎉")
		return
	}

	fmt.Println("Reprovado! 😖")
}
