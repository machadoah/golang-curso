package main

import "fmt"

func inverteSinalV1(num int) int {
	return num * -1
}

func inverteSinalV2(num *int) {
	*num = *num * -1
}

func main() {
	numero01 := 20

	numeroInvertido01 := inverteSinalV1(numero01)

	fmt.Println("Numero01 antes da inversao:", numero01)
	fmt.Println("NumeroInvetido             ", numeroInvertido01)
	fmt.Println("Numero01 após inversao:    ", numero01)
	/*
		nesse caso acima, o numeroInvertido é uma nova variável
		numero é uma variável diferente de numeroInvertido
	*/

	fmt.Println("-------------------------------------")

	numero02 := 40

	fmt.Println("Numero01 antes da inversao:", numero02)
	fmt.Print("NumeroInvertido            ")
	inverteSinalV2(&numero02)

	fmt.Println("\nNumero02 após inversao:    ", numero02)
	/*
		nesse caso acima, o numeroInvertido não é uma nova variável
		numeroInvertido é armazenado no mesmo local de memoria de numero
	*/
}
