package main

import (
	"errors"
	"fmt"
)

func main() {
	// Numeros

	// INTEIROS

	/*
		int8, int16, int32, int64 -> tamanho de bits suportados

		int -> quando o golang usa a arquitetura do computador como base x64, x32
	*/

	var numeroNormal int8 = 100

	fmt.Println(numeroNormal)

	var idade int = 21

	fmt.Println(idade) // assume int64

	/*
		uint -> unsygned int (inteiro sem sinal)
	*/

	var numero uint = 199 // -199 nao funciona
	// # command-line-arguments
	//./tipos-de-dados.go:28:20: cannot use -199 (untyped int constant) as uint value in variable declaration (overflows)

	fmt.Println(numero)

	// alias

	var numero1 rune = 100 // rune = int32
	fmt.Println(numero1)

	var numero2 byte = 100 // byte = int8
	fmt.Println(numero2)

	// REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	numeroReal3 := 123.45
	fmt.Println(numeroReal3)

	// TEXTO

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto"
	fmt.Println(str2)

	// CHAR

	char := 'B' // Valor ascii com aspas simples -> 66
	fmt.Println(char)

	// VALOR ZERO -> VALORES DEFAULTS

	// string
	var texto string
	fmt.Println(texto)

	/*
			string é "" (espaco em branco)
			int ou float é 0 e 0.0
			bool é false
		etc...
	*/

	// BOOLEANO

	var boole bool = true
	fmt.Println(boole)

	// ERRO

	var erro error
	fmt.Println(erro) // <nil>

	// criando um erro!
	var erro2 error = errors.New("Erro Interno")
	fmt.Println(erro2)
}
