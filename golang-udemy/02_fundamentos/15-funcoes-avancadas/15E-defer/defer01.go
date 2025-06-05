package main

import "fmt"

func funcao01() {
	fmt.Println("Executando a funcao 01")
}

func funcao02() {
	fmt.Println("Executando a funcao 02")
}

func alunoAprovado(nota1, nota2 float64) bool {
	fmt.Println("Entrando na função que verifica se o aluno foi aprovado")

	media := (nota1 + nota2) / 2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	defer funcao01() // adia a execução -> executado por ultimo
	funcao02()

	fmt.Println(alunoAprovado(7, 10))
}
