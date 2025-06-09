package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("1a linha")
	fmt.Println("2a linha")
	fmt.Println("3a linha")

	/*
		2a linha
		3a linha
		1a linha

		normalmente defer é usado para fechar conexoes/arquivos etc
	*/
}
