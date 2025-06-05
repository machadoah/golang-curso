package main

import "fmt"

func main() {

	// declaracao explicita
	var nombre string = "Antonio"

	// declarcao implicita - o go identifica o valor (infere)
	sobrenombre := "Machado"

	fmt.Println(nombre)
	fmt.Println(sobrenombre)

	// varias variaveis
	var (
		curso     string = "ADS"
		faculdade string = "Fatec"
	)

	fmt.Println(curso, faculdade)

	// declarando variaveis
	val5, val6 := "a", "b"
	fmt.Println(val5, val6)

	// constante
	const constante string = "Amarelo"
	fmt.Println(constante)

	// invertendo valores de variaveis
	var val7 = "A"
	var val8 = "B"

	fmt.Println(val7, val8) // A B

	val7, val8 = val8, val7

	fmt.Println(val7, val8) // B A
}
