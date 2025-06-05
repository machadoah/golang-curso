package main

import "fmt"

func calculosMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return // pega automaticamente o nome das variaveis que estao no retorno
}

func main() {
	soma, sub := calculosMatematicos(10, 20)
	fmt.Println(soma, sub)
}
