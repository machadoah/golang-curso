package main

import "fmt"

func recExec() {

	if r := recover(); r != nil {
		fmt.Println("recuperando execucao")
	}
}

func alunoIsAprovado(n1, n2 float64) bool {
	defer recExec()

	media := (n1 + n2) / 2

	// numero positivo
	if media >= 0 {

		// aprovado
		if media > 6 {
			return true
			// reprovado
		} else if media < 6 {
			return false
		}
	}

	panic("NOTA MENOR QUE ZERO!") // mata a execucao do programa
}

func main() {
	// fmt.Println(alunoIsAprovado(7, 6))

	fmt.Println(alunoIsAprovado(-3, -4))

	fmt.Println("Pós-execução")
}
