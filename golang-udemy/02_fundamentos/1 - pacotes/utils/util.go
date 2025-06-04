package utils

import "fmt"

// funcoes com letras maiusculas sao publicas

// Escrever msg na tela
func Escrever() {
	fmt.Println("Escrevendo do pacote util")

	// importando a funcao privada do outro arquivo
	escrever2()
}
