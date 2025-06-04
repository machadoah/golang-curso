package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

// funcoes podem ter mais de um retorno!
func calcular(n1, n2 int) (int, int) {
	soma := n1 + n2
	sub := n1 - n2

	return soma, sub
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma) // 30

	var f = func(text string) string {
		fmt.Println("Funcao F -", text)
		return text
	}

	resultado := f("Antonio")
	fmt.Println(resultado)

	rSoma, rSub := calcular(10, 15)
	fmt.Println(rSoma, rSub)

	queroSoma, _ := calcular(rSoma, rSub) // ignora o numero da 2a pos usando o _ underline
	fmt.Println(queroSoma)
}
