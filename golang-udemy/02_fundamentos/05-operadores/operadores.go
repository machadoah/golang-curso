package main

import "fmt"

func main() {
	// OPERADORES ARITMETICOS
	fmt.Println("------------------------\nOPERADORES ARTIMÉTICOS")

	soma := 1 + 2
	sub := 1 - 2
	div := 1 / 2 // 0 pois sao dois numeros sem ponto flutuante
	mult := 1 * 2
	restoDivisao := 10 % 2

	fmt.Println(soma, sub, div, mult, restoDivisao)

	var numero1 int = 10
	var numero2 int = 25

	num1MaisNum2 := numero1 + numero2
	fmt.Println(num1MaisNum2)

	div1 := 1.0 / 2.0
	fmt.Println(div1) // 0.5

	// ATRIBUICAO
	fmt.Println("------------------------\nOPERADORES DE ATRIBUIÇÃO")

	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)

	// OPERAADORES RELACIONAIS
	fmt.Println("------------------------\nOPERADORES RELACIONAIS")

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// OPERADORES LÓGICOS
	fmt.Println("------------------------\nOPERADORES LÓGICOS")

	fmt.Println(true && false && true) // false
	fmt.Println(true || false || true) // true
	fmt.Println(!true)                 // false

	// OPERADORES UNÁRIOS
	fmt.Println("------------------------\nOPERADORES UNÁRIOS")

	numero := 10
	numero++            // numero = numero + 1
	fmt.Println(numero) // 11

	numero--            // numero = numero - 1
	fmt.Println(numero) // 10

	numero += 15        // numero = numero + 15
	fmt.Println(numero) // 25

	numero *= 2         // numero = numero * 2
	fmt.Println(numero) // 50

	// Podemos fazer <numero op= valor> ex: numero *= 2

	// OPERADOR TERNÁRIO
	fmt.Println("------------------------\nOPERADORES TERNÁRIO")
	// Nao existe seguindo a premissa de existirem somente uma forma de fazer algo

	var texto string

	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)
}
