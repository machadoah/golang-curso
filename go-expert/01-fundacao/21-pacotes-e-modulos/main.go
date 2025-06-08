package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/machadoah/go-expert/matematica"
)

func main() {
	// go mod init nome-do-modulo
	s := matematica.Soma(10, 20, 60, 70)

	fmt.Printf("O valor %v\n", s)

	fmt.Println(matematica.Pii)

	carro := matematica.Carro{
		Marca: "FIAT",
	}

	println(carro.Marca)

	fmt.Println(uuid.New())
}
